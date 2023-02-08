#!/usr/local/bin/python3

# Proof of concept / work in progress to capture output from go-edtf-wasm (wasi) binary
# note that the wasi binary doesn't return JSON yet pending updates to whosonfirst/go-edtf
# to provide native marshaling in advance of tinygo supporting encoding/json.
# See cmd/parse-wasi/main.go for details

# https://github.com/wasmerio/wasmer-python/issues/661
# https://stackoverflow.com/questions/24277488/in-python-how-to-capture-the-stdout-from-a-c-shared-library-to-a-variable/29834357#29834357

# > /usr/local/opt/python@3.10/bin/python3.10 ./test-wasmer.py 2022-01~
# 2022-01~

# $> /usr/local/opt/python@3.10/bin/python3.10 ./test-wasmer.py invalid
# Unrecognized EDTF string 'invalid' (Invalid or unsupported EDTF string)

from wasmer import engine, Store, Module, Instance, wasi
from wasmer_compiler_cranelift import Compiler

import os
import sys
import threading
import time

class OutputGrabber(object):
    """
    Class used to grab standard output or another stream.
    """
    escape_char = "\b"

    def __init__(self, stream=None, threaded=False):
        self.origstream = stream
        self.threaded = threaded
        if self.origstream is None:
            self.origstream = sys.stdout
        self.origstreamfd = self.origstream.fileno()
        self.capturedtext = ""
        # Create a pipe so the stream can be captured:
        self.pipe_out, self.pipe_in = os.pipe()

    def __enter__(self):
        self.start()
        return self

    def __exit__(self, type, value, traceback):
        self.stop()

    def start(self):
        """
        Start capturing the stream data.
        """
        self.capturedtext = ""
        # Save a copy of the stream:
        self.streamfd = os.dup(self.origstreamfd)
        # Replace the original stream with our write pipe:
        os.dup2(self.pipe_in, self.origstreamfd)
        if self.threaded:
            # Start thread that will read the stream:
            self.workerThread = threading.Thread(target=self.readOutput)
            self.workerThread.start()
            # Make sure that the thread is running and os.read() has executed:
            time.sleep(0.01)

    def stop(self):
        """
        Stop capturing the stream data and save the text in `capturedtext`.
        """
        # Print the escape character to make the readOutput method stop:
        self.origstream.write(self.escape_char)
        # Flush the stream to make sure all our data goes in before
        # the escape character:
        self.origstream.flush()
        if self.threaded:
            # wait until the thread finishes so we are sure that
            # we have until the last character:
            self.workerThread.join()
        else:
            self.readOutput()
        # Close the pipe:
        os.close(self.pipe_in)
        os.close(self.pipe_out)
        # Restore the original stream:
        os.dup2(self.streamfd, self.origstreamfd)
        # Close the duplicate stream:
        os.close(self.streamfd)

    def readOutput(self):
        """
        Read the stream data (one byte at a time)
        and save the text in `capturedtext`.
        """

        while True:

            char = os.read(self.pipe_out,1).decode(self.origstream.encoding)

            if not char:
                break
            
            if self.escape_char in char:
                 break
            
            self.capturedtext += char


if __name__ == "__main__":

    r = open("../www/wasi/parse.wasm", "rb")

    store = Store(engine.JIT(Compiler))
    
    module = Module(store, r.read())

    wasi_version = wasi.get_version(module, strict=True)

    wasi_env = wasi.StateBuilder('main').  \
        argument(sys.argv[1]). \
        finalize()
    
    import_object = wasi_env.generate_import_object(store, wasi_version)
    instance = Instance(module, import_object)
    
    out = OutputGrabber()
    with out:
        instance.exports._start()
        
    print(out.capturedtext.strip())


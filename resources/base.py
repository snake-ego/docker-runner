import os
import asyncio
from subprocess import STDOUT
from logging import getLogger


class AsyncRunner(object):

    def __init__(self, shell=None, logger=None, command=None):
        self.shell = shell if shell is not None else os.environ['SHELL']
        self.logger = logger if logger is not None else getLogger(self.__class__.__name__)
        self.command = command if isinstance(command, str) else ""

    async def __call__(self, command=None):
        command = command if isinstance(command, str) else self.command

        self.runner = await asyncio.create_subprocess_shell(command, stdout=asyncio.subprocess.PIPE, stderr=STDOUT,
                                                            shell=True, executable=self.shell)
        async for line in self.runner.stdout:
            self.logger.info(str(line, 'utf-8').strip())

#!/usr/bin/env python3

import asyncio
from os import path as op
from argparse import ArgumentParser
from logging import getLogger
from logging.config import dictConfig
from resources.config import ConfigFromJSON
from resources.base import AsyncRunner


def create_tasks(config):
    out = list()
    for logger, task in cfg.get('tasks', list()).items():
        params = cfg.get('defaults', dict()).copy()
        params.update(task)
        params.update({"logger": getLogger(logger)})
        out.append(AsyncRunner(**params))

    return out


def setup_output(logdata):
    return dictConfig(logdata) if isinstance(logdata, dict) else None


def args_parse():
    parser = ArgumentParser()
    parser.add_argument("-c", "--config", type=str, required=False, default='',
                        help="Tasks configuration file")
    return parser.parse_args()


if __name__ == '__main__':
    args = args_parse()
    cfg = ConfigFromJSON(config=args.config) if op.exists(args.config) else ConfigFromJSON()
    setup_output(cfg.get('logging'))

    tasks = create_tasks(cfg)
    loop = asyncio.get_event_loop()

    [asyncio.ensure_future(t()) for t in tasks]
    try:
        loop.run_forever()
    except (KeyboardInterrupt, SystemExit):
        for t in asyncio.Task.all_tasks():
            t.cancel()

        loop.stop()
        loop.close()

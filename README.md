# Description

Simple multiporcesses runner for docker. 
Uses python 3.5 asyncio library

# Configuration

```javascript
{
    "tasks": {
        "name": {
            "command": "exit 0"
        }
    },

    "defaults": {
        "shell": "/bin/sh"
    },

    "logging": {
        "version": 1,
        "disable_existing_loggers": false,

        "formatters": {
            "c_short": {
                "format": "%(message)s",
                "datefmt": "%Y-%m-%d %H:%M:%S"
            }
        },

        "handlers": {
            "default_stream": {
                "class": "logging.StreamHandler",
                "level": "INFO",
                "formatter": "c_short"
            }
        },

        "loggers": {
            "": {
                "handlers": ["default_stream"],
                "level": "INFO",
                "propagate": true
            }
        }
    }
}
```

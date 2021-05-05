# go-tg-wol
Telegram Bot for wake on lan

# Run
The file `config.json` must be created near executable file.

**Config.json example**
```json
{
    "token": "CHANGEME",
    "users": [
        {
            "name": "Patrick",
            "phone": "+799999999",
            "mac_address": "01:02:03:04:05:07"
        }
    ],
    "log_file": "bot.log"
}
```

Then just run `./bot` or `.\bot.exe` depends on your OS.

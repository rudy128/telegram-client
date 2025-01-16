# Telegram Media Bot

This project is a Telegram client bot built using the [gotd](https://github.com/gotd/td) library in Go. The bot listens to messages in a specific chat or channel, downloads media files (photos, videos, etc.), and processes them. The bot is designed to handle concurrency safely and ensures that all downloads complete successfully.

## Features

- Connects to Telegram via MTProto API.
- Listens for new messages in a specified chat or channel.
- Downloads media files from messages and stores them locally.
- Ensures file integrity and prevents corruption during downloads.
- Sends media or text messages to specific chats.
- Graceful shutdown handling for active downloads.

## Prerequisites

1. **Telegram API Credentials**:
   - Obtain `APP_ID` and `APP_HASH` from [my.telegram.org](https://my.telegram.org/).

2. **Go Environment**:
   - Install Go (1.18 or later is recommended).

3. **Environment Variables**:
   - Set the following environment variables:
     ```bash
     export TG_PHONE="+1234567890"  # Your phone number in international format
     export APP_ID="your_app_id"    # Telegram App ID
     export APP_HASH="your_app_hash"  # Telegram App Hash
     export CHAT_ID="123456789"     # Target Chat ID for messages
     ```

4. **Directories**:
   - Ensure the directory for storing session and downloaded files exists:
     ```bash
     mkdir -p session downloads
     ```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/telegram-media-bot.git
   cd telegram-media-bot
   ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```
3. Build the project:
    ```bash
    go build -o telegram-media-bot .
    ```

## Usage
Run the bot with the following command:
```bash
./telegram-media-bot
```
The bot will:
- Authenticate using the phone number provided.
- Listen for messages in the target chat specified by ```CHAT_ID```.
- Automatically download media files and save them to the ```downloads``` directory.


## Configuration
### Environment Variables
You can configure the bot using environment variables:
| Variable | Description |
|----------|----------|
| ```TG_PHONE``` | Your phone number in international format. |
| ```APP_ID``` | Telegram App ID obtained from Telegram. |
| ```APP_HASH``` | Telegram App Hash obtained from Telegram. |
| ```CHAT_ID``` | The chat/channel ID to monitor and send messages to. |
| ```RECEPTIONIST_ID``` | The user ID to whom messages are sent. |


## Directories
The bot uses the following directories:

- ```session/```: Stores session data for authentication persistence.
- ```downloads/```: Saves downloaded media files.


## Project Structure

/project-root
├── main.go                      # Entry point for the bot
├── src/                         # Source files
│   ├── channel_specific.go      # Handles specific channel logic
│   ├── check_create_folder.go   # Ensures necessary folders exist
│   ├── download_media.go        # Handles media downloads
│   ├── empty_dir.go             # Empties a directory's contents
│   ├── extract_user.go          # Extracts numerical user/channel IDs
│   ├── file_location.go         # Extracts file location for media
│   ├── global.go                # Global variables and sync management
│   ├── random_file.go           # Picks a random file from a directory
│   ├── run.go                   # Initializes and runs the bot
│   ├── send_message.go          # Sends photos and captions to a user
│   ├── session.go               # Creates Session for the user
├── downloads/                   # Directory for downloaded media files
├── session/                     # Directory for storing session data
├── go.mod                       # Defines dependencies and module metadata.
├── go.sum                       # Ensures dependency integrity with cryptographic hashes.
├── Dockerfile.armv8             # A multi-stage Dockerfile to build and run a Go app on ARM64.
├── .env                         # A file for storing environment variables.
└── README.md                    # Project documentation

## Key Functions

### ```moneyChannel```
Handles incoming messages in the monitored chat. Identifies relevant messages, initiates media downloads, and triggers the sending of messages with captions.

### ```download_media```
Downloads media from Telegram messages and saves them with unique names to the downloads directory.

### ```sendMessage```
Sends a photo with a caption to a specified user or chat. Cleans up the downloads directory after sending.

### ```checkAndCreateFolder```
Ensures that required directories (session and downloads) exist before running the bot.

### ```extractUserID```
Extracts the numerical ID from Telegram's peer objects (PeerChannel, PeerChat, PeerUser).

### ```getInputFileLocation```
Extracts the input file location for a media file in a message, enabling the download process.

### ```getRandomFileFromDir```
Selects a random file from the specified directory, used to pick a media file to send.

### ```EmptyDir```
Clears all files from a directory without removing the directory itself.

### ```Run```
Initializes and runs the Telegram client, authenticates the user, and handles message updates.

## Logging and Debugging
Logs are written to the console for tracking events such as:

- Successful connections.
- Media downloads.
- Errors during processing.

You can extend the logging functionality using a tool like ```zap``` for structured logs.

## Graceful Shutdown
The bot listens for termination signals (e.g., ```Ctrl+C```) and waits for all active downloads to finish before exiting. This ensures no files are corrupted during abrupt exits.

## Future Improvements
Here are some ideas for enhancing the bot:

- Support for more media types, such as documents or voice notes.
- Dynamic configuration updates without restarting the bot.
- Improved retry logic for failed downloads.
- Integration with cloud storage for media backup.

# Contribution
Contributions are welcome! If you encounter a bug or have suggestions for improvements, feel free to open an issue or submit a pull request.
To contribute:

1. Fork the repository.
2. Create a new branch (```git checkout -b feature-name```).
3. Make your changes and commit them (```git commit -m "Add feature-name"```).
4. Push to the branch (```git push origin feature-name```).
5. Open a pull request.

# Acknowledgements
- gotd: Telegram MTProto implementation for Go.
- my.telegram.org: Telegram API registration portal.
- Go developers and contributors for providing an excellent language and ecosystem.
from telethon import TelegramClient, events
import sys

if len(sys.argv) != 5:
    sys.exit(1)
    
phone_number = sys.argv[1]
api_id = sys.argv[2]
api_hash = sys.argv[3]
userID = sys.argv[4]

client = TelegramClient('session_name', api_id, api_hash)

async def get_dialogs_and_user_ids():
    # Connect to the Telegram client
    await client.start(phone_number)

    try:
        # Get all dialogs (chats)
        dialogs = await client.get_dialogs()

        # Loop through the dialogs and extract userIDs (PeerUser)
        for dialog in dialogs:
            if dialog.is_user:
                # Check if it's a user (not a group or channel)
                if userID == str(dialog.entity.id):
                    # Print the username as output for Go to capture
                    print(dialog.entity.username)
                    return

        # If no matching user found
        print("User not found")
    finally:
        # Disconnect the client after the operation
        await client.disconnect()
            



import asyncio
asyncio.run(get_dialogs_and_user_ids())
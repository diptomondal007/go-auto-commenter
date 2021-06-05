#! /bin/bash

TARGET=/usr/local/bin/autocomment
MESSAGE_START="Removing Go Auto Commenter"
MESSAGE_END="√ Go Auto Commenter removed"

echo "$MESSAGE_START"
rm $TARGET
echo "$MESSAGE_END"
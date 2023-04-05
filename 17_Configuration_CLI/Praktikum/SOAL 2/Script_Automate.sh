#!/bin/bash

# membuat variabel dinamis
main_folder=$1
facebook_username=$2
linkedin_username=$3

# Membuat direktori main
mkdir -p "$main_folder"

# membuat direktori aboutme dan membuat subflder personal dan profesional didalamnya
mkdir -p "$main_folder/about_me/personal"
mkdir -p "$main_folder/about_me/professional"

# membuat my friend folder
mkdir -p "$main_folder/my_friends"

# membuat folder my sistem info
mkdir -p "$main_folder/my_system_info"

# membuat file facebook.txt
echo "https://www.facebook.com/$facebook_username" > "$main_folder/about_me/personal/facebook.txt"

# membuat file linkedin.txt
echo "https://www.linkedin.com/in/$linkedin_username" > "$main_folder/about_me/professional/linkedin.txt"

# membuat file list_of_my_friends.txt
curl -o "$main_folder/my_friends/list_of_my_friends.txt" https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt

# Membuat about_this_laptop.txt
echo "My username: $(whoami)" > "$main_folder/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$main_folder/my_system_info/about_this_laptop.txt"

# Membuat file  internet_connection.txt
ping -c 3 google.com > "$main_folder/my_system_info/internet_connection.txt"

# jika semua proses sukses semua maka print pesan
echo "Automation process completed successfully!"
#!/usr/bin/env bash

# membuat variable
directory_name="$1"
facebook_name="$2"
linkedin_name="$3"
current_date="$(date "+%a %b %d %T %Z %Y")"
username="$(whoami)"
host="$(uname -a)"

# membuat main directory yang merupakan gabungan dari argumen pertama dan current_date
main_directory="$directory_name at $current_date"

# membuat main directory 
mkdir -p "$main_directory"

# membuat struktur directory 
mkdir -p "$main_directory/about_me/personal"
mkdir -p "$main_directory/about_me/professional"
mkdir -p "$main_directory/my_friends"
mkdir -p "$main_directory/my_system_info"

# membuat file
touch "$main_directory/about_me/personal/facebook.txt"
touch "$main_directory/about_me/professional/linkedin.txt"
touch "$main_directory/my_friends/list_of_my_friends.txt"
touch "$main_directory/my_system_info/about_this_laptop.txt"
touch "$main_directory/my_system_info/internet_connection.txt"

# menulis file facebook.txt dan linkedin.txt yang berisikan url sesuai argumen kedua dan ketiga
echo "https://www.facebook.com/$facebook_name" > "$main_directory/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$linkedin_name" > "$main_directory/about_me/professional/linkedin.txt"

# menulis file list_of_my_friends.txt menggunakan curl
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$main_directory/my_friends/list_of_my_friends.txt"

# menulis file about_this_laptop.txt berisikan nama user dan uname -a 
echo "my username : $username" > "$main_directory/my_system_info/about_this_laptop.txt"
echo "with host : $host" > "$main_directory/my_system_info/about_this_laptop.txt"

# menulis file internet_connection.txt berisikan hasil keluaran ping ke google.com sebanyak 3 kali
ping -c 3 google.com > "$main_directory/my_system_info/internet_connection.txt"

# menampilkan struktur directory
tree "$main_directory"
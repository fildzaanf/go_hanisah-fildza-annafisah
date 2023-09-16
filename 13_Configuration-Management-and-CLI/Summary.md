# Summary

## (13) Configuration Management dan Command Line Interface (CLI)

### Command Line Interface (CLI)
- Command Line Interface (CLI) adalah antarmuka berbasis teks yang memungkinkan pengguna berinteraksi dengan sistem melalui baris perintah secara efektif dan efisien
- Shell = CLI for OS services
    - UNIX Shell
    - Command Prompt (MDSOS)

- Other app CLI 
- Python REPL
- MySQL CLI Client
- Mongo Shell
- redis-CLI

### Why using Command Line 
- granular control dari OS atau aplikasi
- manajemen lebih cepat dari sejumlah besar sistem operasi
- kemampuan untuk menyimpan script untuk mengotomatisasikan tugas-tugas rutin
- menghubungkan pengetahuan untuk membantu pemecahan masalah, seperti masalah koneksi jaringan

### Intro to UNIX Shell 
- Example shell : 
```sh
me@linuxbox:~$

```
keterangan 
    - me : username
    - linuxbox : hostname
    - ~ : current path (home)
    - $ : shell for normal username
    - # : shell for root user

### Normal User vs Root User
- Normal User = allowed to manupulating /home/username dir only
- Root User = allowed to manupulating /dir (all directory)
- Normal User + Sudoers = allowed to act as root temporary

### UNIX Shell Most Popular Command
#### Directory Commands
- `pwd`: Print working directory
- `ls`: List files and directories
- `mkdir`: Create a new directory
- `cd`: Change directory
- `rm`: Remove files or directories
- `cp`: Copy files or directories
- `mv`: Move or rename files and directories
- `ln`: Create links

#### File Commands
- Create: `touch`
- View: `head`, `cat`, `tail`, `less`
- Editor: `vim`, `nano`
- Permission: `chown`, `chmod`
- Different: `diff`

#### Network Commands
- `ping`: Send ICMP echo requests to a host
- `ssh`: Secure Shell remote login
- `netstat`: Network statistics
- `nmap`: Network mapper
- `ip addr` (ifconfig): Show network interfaces and configuration
- `wget`: Download files from the web
- `curl`: Transfer data with URLs
- `telnet`: Connect to remote hosts using the Telnet protocol
- `netcat`: Network utility for reading and writing to network connections

#### Utility Commands
- `man`: Display manual pages
- `env`: Display environment variables
- `echo`: Display text
- `date`: Display  the system date and time
- `which`: Locate a command
- `watch`: Execute a program periodically
- `sudo`: Execute a command as another user
- `history`: Command history
- `grep`: Search text using patterns
- `locate`: Search for files and directories
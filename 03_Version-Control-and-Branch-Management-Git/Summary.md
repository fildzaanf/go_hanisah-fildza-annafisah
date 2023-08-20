# Summary 

## (03) Version Control and Branch Management (Git)

### Version Control Systems (VCS) and Repo Hosting Services
- Version Control Systems adalah tools untuk mengelola perubahan dan konfigurasi pada aplikasi termasuk source code
- Version Control Systems dapat melacak dan merekam history perubahan yang dilakukan sehingga dapat mempermudah perbaikan bug
- Version Control Systems juga dapat melakukan pemulihan (backup) dan pembatalan perubahan
- Version Control Systems merupakan tools penting developer terutama developer yang berkolaborasi dalam membangun aplikasi
- Version Control Systems yang dapat digunakan oleh developer diantaranya adalah Git
- Repo Hosting Services adalah platform yang dapat melacak perubahan, mengelola branch, menggabungkan (merge) kode, dan menyimpan serta mengelola repository Git secara online yang menggunakan penyimpanan berbasis cloud sehingga mempermudah developer dalam berkolaborasi membuat aplikasi
- Repo Hosting Services yang dapat digunakan oleh developer diantaranya sebagai berikut.
    - GitHub
    - GitLab
    - BitBucket

### Git

#### Setting up

```
# git config
$ git config --global user.name “John Done”
$ git config --global user.email “johndoe@email.com”

# start with init
$ git init
$ git remote add <remote_name> <remote_repo_url>
$ git push -u <remote_name> <local_branch_name>

# start with existing project, start working on the project
$ git clone ssh://john@example.com/path/to/my-project.git
$ cd my-project 

```

#### Saving Changes

terbagi menjadi 3 fase yaitu working directory, staging area, dan repository

```
$ git status

# memindahkan ke staging area
$ git add <directory>
$ git add hello.py
$ git add .

# memindahkan ke repository
$ git commit -m “add config file”

```
#### Git Diff and Stash

```
# git diff
# change file
# add staging area
$ git diff --staged

# stashing your work
$ git stash	

# re-applying your stashed changes
$ git stash apply

```
####  Git Log, Checkout, Reset

```
# viewing an old revision
$ git log --oneline

$ git checkout a1e8fb5

$ git reset a1e8fb5 --soft

```
#### Syncing

```
# git remote
$ git remote -v
$ git remote add origin http://dev.example.com/john.git

# fetch and pull
$ git fetch 
$ git pull origin master

# push
$ git push origin master
$ git push origin feature/login-user

```

#### Branching

```
# show all branch
$ git branch --list

# create a new branch called <branch>
$ git branch <branch>

# force delete the specified branch
$ git branch -D <branch>

# list remote branch
$ git branch -a

# Start a new feature
$ git checkout -b new-feature master
# Edit some files
$ git add <file>
$ git commit -m "Start a feature"
# Edit some files
$ git add <file>
$ git commit -m "Finish a feature"
# Merge in the new-feature branch
$ git checkout master
$ git merge new-feature
$ git branch -d new-feature

```

#### Perintah dan Fungsi pada Git

- Git Repository    : Media penyimpanan file project di dalam Git server
- Git Branch        : Percabangan untuk versi baru dari project repository
- Git Fork          : Penyalinan dari repository lain dan menyimpannya di repository Git sendiri
- Git Clone         : Mengambil repository server dan menyimpannya di direktori lokal
- Git Commit        : Cuplikan perubahan (snapshot) dari repository pada waktu tertentu
- Git Push          : Mengirim hasil dari perubahan file yang dilakukan ke repository server
- Git Pull          : Mengambil source code dari repository server ke lokal
- Git Checkout      : Melakukan perpindahan dari commit satu ke commit yang dituju
- Git Reset         : Mengembalikan keadaan commit ke dalam keadaan sebelum terjadi perubahan sesuai dengan commit yang dituju. Git reset akan menghapus beberapa riwayat commit sesudahnya
- Git Revert        : Mengembalikan keadaan commit sebelum terjadi suatu perubahan sesuai dengan tujuan commit yang dituju. Git Revert tidak akan menghilangkan riwayat commit sesudahnya



### Workflow Collaboration
- Buat beberapa branch untuk kolaborasi diantaranya sebagai berikut.
    - Master  
    - Development 
    - Feature 
- Cara mengoptimalkan kolaborasi dalam development
    - Buat branch Master dari branch Development
        ```
        $ (master) git branch development
        $ (master) git checkout development
        
        ```
    - Hindari direct edit ke branch Development
        ```
        $ (development) git branch feature1
        $ (development) git checkout feature1

        ```
    - Merge branch Feature hanya ke branch Development
        
        ```
        $ (feature1) git checkout development
        $ (development) git merge feature1

        ```
    - Merge branch Development ke branch Master jika development sudah selesai
        ```
        $ (master) git merge development

        ```







 
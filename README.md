# Project 2: Shell Builtins

## Description

In this project, we have a shell that inputs the following commands:

-ls

-cd

-mkdir

-rmdir

-mv

-echo

-env

## Brief descriptions of commands

### -ls-

This command is used to list all files and directories in the present directory.
Use: ls

### -cd-

This command is used to change directories.
Use: cd [takes user back to the home directory]
     cd <Directory_name> [takes user into the directory with that name, as long as it is accessible from within the present working directory]
     cd .. [takes user back by one directory (parent directory)]

### -mkdir-

This command is used to make new directories.
Use: mkdir <Directory_name>

### -rmdir-

This command is used to remove directories in the present working directory.
Use: rmdir <Directory_name>

### -mv-

This command is used to move a file from one directory to another (source, destination). This command requires the full path.
Use: mv <source_path> <destination_path>

### -echo-

This command is used to echo whatever is typed after the command.
Use: echo <statement>

### -env-

This command is used to display the user's current environment or run a specified command in a changed environment.
Use: env

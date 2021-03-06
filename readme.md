[![GoDoc](https://godoc.org/github.com/hithacker/copymy?status.svg)](https://godoc.org/github.com/hithacker/copymy)

Copy anything that you need to paste frequently to clipboard by using this program. For example when applying for jobs you often need to enter your linkedin url or github url and you may not always remember it so you have to waste time looking for it in your browser history or even worse you need to open the linkedin website and copy the url. This command will allow you to make this process much faster. Most of us programmers always have the terminal open or we can bring it up by a shortcut so you just need to enter "copymy name-of-the-thing-you-want-to-copy" on terminal to copy the thing we want to clipboard.

##Configuration
To make this work you will need to create a file named copymy_config.toml in your user's home directory e.g. for me my home directory on OSX is /Users/hithacker. It might be something different for you depending on your username and the operating system you are using. Add the things you often need to copy in this file in key value format. You can take a look at my config file as an example https://github.com/hithacker/copymy/blob/master/copymy_config.toml

##Usage
Just type the command name followed by the name of the thing you want to copy
    
    $ copymy linkedin
    Copied https://www.linkedin.com/in/hirenthacker to clipboard
    $ copymy github
    Copied https://github.com/hithacker to clipboard



Let me know if you have any questions by mailing me on hithacker@gmail.com or by raising an issue.

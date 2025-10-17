> [!WARNING]  
> Gnome 49 has broken the old gnome-screenshot tool (which was already deprecated). This means this hack will not work on Gnome 49+. See [issue #1](https://github.com/teoulas/gnome-screenshot2swappy/issues/1). If you're looking for an alternative, consider [Gradia](https://gradia.alexandervanhee.be/).
> 
# What

A script to assign to a keyboard shortcut (e.g. super+shift+s like Windows), which uses gnome-screenshot to select an area, then passes it to [swappy](https://github.com/jtheoof/swappy) for annotations and then opens nautilus with the file preselected (if you save it to disk).

![Demo screencast](img/demo-screencast.gif)

# How

A bash and go script as glue between the apps.

You need to add the custom shortcut to the custom keyboard shortcuts in Gnome settings. See:

![Swappy Screenshot](img/gnome-keyboard-shortcut.png)

# Why

I wanted a quick and simple way to take a screenshot, annotate it and be able to copy/paste (or drag'n'drop).

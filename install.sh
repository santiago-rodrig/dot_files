#!/bin/bash

# declare main variables

VIMFILE="$HOME/.vimrc"
TMUXFILE="$HOME/.tmux.conf"
LIGHT_GREEN='\e[92m'
LIGHT_CYAN='\e[96m'
RESET_COLORS='\e[0m'
INVERTED_COLORS='\e[7m'
RESET_INVERTED_COLORS='\e[27m'

# create a temporary alias for colored output

alias echo='echo -e'

# define function to reset the colors after coloring output

function resetColors () {
  echo -n $RESET_INVERTED_COLORS$RESET_COLORS
  echo $RESET_INVERTED_COLORS$RESET_COLORS
}

# define function to color the output of an informational message

function beginInfo () {
  echo -n $LIGHT_CYAN$1
  resetColors
  echo : $INVERTED_COLORS
}

# define function to color the output of a creation message

function beginCreate () {
  echo $LIGHT_GREEN$1
  resetColors
  echo : $INVERTED_COLORS
}


# move the configuration files to their corresponding place

beginInfo 'MOVE'
echo "\ttmux.conf -> $HOME/.tmux.conf"
echo "\tvimrc -> $HOME/.vimrc"
resetColors
mv vimrc $VIMFILE
mv tmux.conf $TMUXFILE

# this function expects the intended path where to place the plugin
# $1 is something in the form ./colors/start

function checkDir () {
  if [ ! -d $1 ]; then
    beginCreate 'CREATE'
    echo "\t$1"
    resetColors
    mkdir -p $1
  fi
}

# once inside $HOME/.vim/pack this function checks for repositories

function checkRepo () {
  if [ -d $1 ] && ls "$1/.git"; then
    beginCreate 'UPDATE'
    echo "\tupdating $1"
    resetColors
    cd $1 && git pull
  elif [ -d $1 ] && ! ls "$1/.git"; then
    beginCreate 'CREATE'
    echo "\t$1"
    resetColors
    rm -rf $1 && git clone https://github.com/$2.git $1
  elif ! [ -d $1 ]; then
    beginCreate 'CREATE'
    echo "\t$1"
    resetColors
    git clone https://github.com/$2.git $1
  fi
}

# check for the existence of $HOME/.vim/pack

checkDirs $HOME/.vim/pack

# change directory to $HOME/.vim/pack

beginInfo 'CHANGE DIRECTORY'
resetColors
echo "\t$(pwd) -> $HOME/.vim/pack"
cd $HOME/.vim/pack

# check for the existence of each plugin category

function parsePlugin () {
  TAIL=`echo -n $1 | cut -d '/' -f 2`

  checkRepo $2/$TAIL $1
}

# loop through the categories and get the plugins for each in turn

for $CATEGORY in 'colors' 'editing' 'git' 'html' 'markdown' 'navigation' \
  'ruby' 'status-bar' 'textobjects' 'tmux'; do
  FULL_DIR="./$CATEGORY/start/"

  checkDir $FULL_DIR

  case $CATEGORY in
    'colors' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'morhetz/gruvbox'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'editing' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'tpope/vim-commentary' 'tpope/vim-surround' \
        'christoomey/vim-system-copy'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'git' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'Xuyuanp/nerdtree-git-plugin' 'tpope/vim-fugitive' \
        'airblade/vim-gitgutter'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'html' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'mattn/emmet-vim' 'AndrewRadev/tagalong.vim'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'markdown' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'iamcco/markdown-preview.nvim'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'navigation' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'ctrlpvim/ctrlp.vim' 'preservim/nerdtree'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'ruby' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'tpope/vim-bundler' 'tpope/vim-rails'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'status-bar' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'vim-airline/vim-airline' \
        'vim-airline/vim-airline-themes'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'textobjects' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'kana/vim-textobj-user' 'vim-textobj-entire'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    'tmux' )
      beginInfo 'INFO'
      echo "\tgetting $CATEGORY category plugins"
      resetColors

      for $PLUGIN in 'edkolev/tmuxline.vim'; do
        beginInfo 'INFO'
        echo "\tparsing $PLUGIN"
        resetColors
        parsePlugin $PLUGIN $FULL_DIR
      done;;
    *) beginInfo 'INFO'; echo "\tfinished plugins' gathering";;
  esac
done

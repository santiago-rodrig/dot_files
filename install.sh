VIMFILE="$HOME/.vimrc"
TMUXFILE="$HOME/.tmux.conf"
LIGHT_GREEN='\e[92m'
LIGHT_CYAN='\e[96m'
RESET_COLORS='\e[0m'
INVERTED_COLORS='\e[7m'
RESET_INVERTED_COLORS='\e[27m'

alias echo='echo -e'

echo "${LIGHT_CYAN}MOVE: $INVERTED_COLORS"
echo "\ttmux.conf -> $HOME/.tmux.conf"
echo "\tvimrc -> $HOME/.vimrc$RESET_INVERTED_COLORS$RESET_COLORS"
mv vimrc $VIMFILE
mv tmux.conf $TMUXFILE

# this function clones the repo user and project supplied in $1

function cloneRepo () {
  git clone https://github.com/$1 $2
}

# this function expects the intended path where to place the plugin
# $1 is something in the form ./colors/start

function checkDirs () {
  if [ ! -d $1 ]; then
    mkdir -p $1
    echo "${LIGHT_GREEN}CREATE DIRS: $INVERTED_COLORS"
    echo "\t$1$RESET_INVERTED_COLORS$RESET_COLORS"
  fi
}

# check for the existence of $HOME/.vim/pack

checkDirs "$HOME/.vim/pack"

# change directory to $HOME/.vim/pack

echo "${LIGHT_CYAN}CHANGE DIRECTORY: $INVERTED_COLORS"
colors
editing
git
html
markdown
navigation
ruby
status-bar
textobjects
tmux

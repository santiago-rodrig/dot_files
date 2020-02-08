mkdir -p "~/.vim/autoload" "~/.vim/bundle" && \
curl -LSso "~/.vim/autoload/pathogen.vim" "https://tpo.pe/pathogen.vim"
cd "~/.vim/bundle"
git clone https://github.com/ctrlpvim/ctrlp.vim.git ctrlp
git clone https://github.com/morhetz/gruvbox.git gruvbox
git clone https://github.com/tpope/vim-commentary.git vim-commentary
git clone https://github.com/tpope/vim-surround.git vim-surround
git clone https://github.com/christoomey/vim-system-copy.git vim-system-copy
git clone https://github.com/kana/vim-textobj-user.git vim-textobj-user
git clone https://github.com/kana/vim-textobj-entire.git vim-textobj-entire

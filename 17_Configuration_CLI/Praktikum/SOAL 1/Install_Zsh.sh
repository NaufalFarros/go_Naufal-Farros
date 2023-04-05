
# update package
sudo apt-get update
sudo apt upgrade
#Install prerequisite packages (ZSH, powerline & powerline fonts) 
sudo apt install zsh
sudo apt-get install powerline fonts-powerline
# Clone the Oh My Zsh Respo 
git clone https://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh
# Create a New ZSH configuration file 
cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc
# Set up a Fancy theme for your Terminal - Open .zshrc File using nano editor 
nano .zshrc
# Find the line ZSH_THEME="robbyrussell" replace robbyrussell with agnoster theme in .zshrc File (CTRL + X & Enter to Save) 
ZSH_THEME="agnoster"
# Change your Default Shell 
chsh -s /bin/zsh
# Update & Uninstallation oh-my-zsh Visit - https://github.com/robbyrussell/oh-my-zsh#manual-updates
cd .oh-my-zsh
upgrade_oh_my_zsh
#Want Syntax Highlighting? install ZSH Syntax Highlighting for Oh My Zsh 
# Clone the ZSH Syntax Highlighting
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git "$HOME/.zsh-syntax-highlighting" --depth 1
# Add syntax-highlighting in .zshrc Configuration 
echo "source $HOME/.zsh-syntax-highlighting/zsh-syntax-highlighting.zsh" >> "$HOME/.zshrc"
#Revert Back to Default Shell
chsh -s /bin/bash
# logout / reboot your system and login again

# Prepare ansible with pyenv

# Install all essential packages

sudo apt update

sudo apt install -y \
build-essential \
curl \
git \
make \
gcc \
libssl-dev \
zlib1g-dev \
libbz2-dev \
libreadline-dev \
libsqlite3-dev \
wget \
llvm \
libncurses5-dev \
libncursesw5-dev \
xz-utils \
tk-dev \
libffi-dev \
liblzma-dev \
python3-openssl

# Install python 3.12

pyenv install 3.12

# Create virtualenv name ansible

pyenv virtualenv 3.12.13 ansible

# install ansible

pip install --upgrade pip
pip install ansible

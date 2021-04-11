set -e

cd _generate

# Regenerate the README file
go get -v 
go run .

cd ..


git config --global user.name 'xarantolus'
git config --global user.email '32465636+xarantolus@users.noreply.github.com'

git pull

# Add the file
git add README.md

# This fails if there are no changes
git commit -m"Automatic update"

git push

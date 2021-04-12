set -e

git pull

cd _generate

# Regenerate the README file
go get -v 
go run .

cd ..


git config --global user.name 'github-actions'
git config --global user.email '41898282+github-actions[bot]@users.noreply.github.com'

# Add the file
git add README.md

# This fails if there are no changes
git commit -m"Automatic update"

git push

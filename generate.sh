set -e

cd _generate

# Regenerate the README file
go get -v 
go run .

cd ..

# Add the file
git add README.md

# This fails if there are no changes
git commit -m"Automatic update"

git push

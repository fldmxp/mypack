git init
git add .
git commit -m "test6"
git remote rm origin
git remote add origin git@github.com:fldmxp/mypack.git
git config --global credential.helper store
git push origin master
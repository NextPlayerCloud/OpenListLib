# git remote add openlist https://github.com/OpenListTeam/OpenList.git
git fetch openlist main --tags
git merge openlist/main --allow-unrelated-histories
# git tag v2025.07.04
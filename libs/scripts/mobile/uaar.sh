mvn deploy:deploy-file \
-Dmaven.test.skip=true -Djacoco.skip=true -PCI \
-Dfile=./mobile.aar \
-DgroupId=com.github.openlistapp \
-DartifactId=AListMobile \
-Dversion=1.0.0 \
-Dpackaging=aar \
-DrepositoryId=github \
-Durl=https://maven.pkg.github.com/AListWeb/AListLib \
-X
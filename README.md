# patchman-engine deps backup
Created to backup all dependency source codes needed to build [patchman-engine](https://github.com/RedHatInsights/patchman-engine) project.

Use it in case you are not able to build patchman-engine because some
dependency repository dissapeared from its original location (typically GitHub).

## Update backup
You should backup each change of go.{mod,sum} in patchman-engine. Ensure your
vendor folder to be synced with the current go.* version:
~~~bash
cd patchman-engine
go mod vendor # Download current dependencies sources
cd ../patchman-engine-dep
cp ../patchman-engine/go.* . # updated go.{mod,sum} deps configs
cp -r ../patchman-engine/vendor . # update all deps sources
~~~

Note: Of course it would be good to automate this somehow ;-). E.g. using github action job in patchman-engine repo. It shouldn't be too hard.


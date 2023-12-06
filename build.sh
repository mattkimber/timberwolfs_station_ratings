#!/bin/bash
ls voxel/*.vox | xargs ../gorender/renderobject.exe -8 -r -s 2,1 -u -p 

./cargosprites.exe
cat pnml/header.pnml pnml/spritesets.pnml pnml/economies.pnml > cargo_rating.nml

../nml/nmlc.exe -c cargo_rating.nml
mv cargo_rating.grf timberwolfs_ratings
cp grf_readme/* timberwolfs_ratings

tar -c timberwolfs_ratings > timberwolfs_ratings.tar
rm timberwolfs_ratings/*
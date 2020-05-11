mongoimport -u root -p 123456 -d victoriadb -c admins -h 0.0.0.0:27017 --authenticationDatabase=admin victoriadb_admins.dat
mongoimport -u root -p 123456 -d victoriadb -c employers -h 0.0.0.0:27017 --authenticationDatabase=admin victoriadb_employers.dat
mongoimport -u root -p 123456 -d victoriadb -c responses -h 0.0.0.0:27017 --authenticationDatabase=admin victoriadb_responses.dat
mongoimport -u root -p 123456 -d victoriadb -c videos -h 0.0.0.0:27017 --authenticationDatabase=admin victoriadb_videos.dat
mongoimport -u root -p 123456 -d victoriadb -c candidates -h 0.0.0.0:27017 --authenticationDatabase=admin victoriadb_candidates.dat
mongoimport -u root -p 123456 -d victoriadb -c interviews -h 0.0.0.0:27017 --authenticationDatabase=admin victoriadb_interviews.dat
mongoimport -u root -p 123456 -d victoriadb -c system.indexes -h 0.0.0.0:27017 --authenticationDatabase=admin victoriadb_system.indexes.dat



mongoimport  -d victoriadb -c admins -h 0.0.0.0:27017  victoriadb_admins.dat
mongoimport  -d victoriadb -c employers -h 0.0.0.0:27017  victoriadb_employers.dat
mongoimport  -d victoriadb -c responses -h 0.0.0.0:27017  victoriadb_responses.dat
mongoimport  -d victoriadb -c videos -h 0.0.0.0:27017  victoriadb_videos.dat
mongoimport  -d victoriadb -c candidates -h 0.0.0.0:27017  victoriadb_candidates.dat
mongoimport  -d victoriadb -c interviews -h 0.0.0.0:27017  victoriadb_interviews.dat
mongoimport  -d victoriadb -c system.indexes -h 0.0.0.0:27017  victoriadb_system.indexes.dat


mongoimport  -d victoriadb -c admins -h 127.0.0.1:27017  victoriadb_admins.dat
mongoimport  -d victoriadb -c employers -h 127.0.0.1:27017  victoriadb_employers.dat
mongoimport  -d victoriadb -c responses -h 127.0.0.1:27017  victoriadb_responses.dat
mongoimport  -d victoriadb -c videos -h 127.0.0.1:27017  victoriadb_videos.dat
mongoimport  -d victoriadb -c candidates -h 127.0.0.1:27017  victoriadb_candidates.dat
mongoimport  -d victoriadb -c interviews -h 127.0.0.1:27017  victoriadb_interviews.dat
mongoimport  -d victoriadb -c system.indexes -h 127.0.0.1:27017  victoriadb_system.indexes.dat

/*DROP DATABASE fizzDB*/
/*CREATE DATABASE fizzDB;*/
USE fizzDB;

CREATE OR REPLACE TABLE stats(
    id VARCHAR(150),
    queries INT, 
    PRIMARY KEY(id)
);
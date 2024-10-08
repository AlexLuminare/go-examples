PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE Categories (
    Id INTEGER NOT NULL PRIMARY KEY, 
    Name TEXT
);
INSERT INTO Categories VALUES(1,'Watersports');
INSERT INTO Categories VALUES(2,'Soccer');
CREATE TABLE Products (
    Id INTEGER NOT NULL PRIMARY KEY, 
    Name TEXT,
    Category INTEGER,
    Price decimal(8, 2),
    CONSTRAINT CatRef FOREIGN KEY(Category) REFERENCES Categories (Id)
);
INSERT INTO Products VALUES(1,'Kayak',1,279);
INSERT INTO Products VALUES(2,'Lifejacket',1,48.95000000000000284);
INSERT INTO Products VALUES(3,'Soccer Ball',2,19.5);
INSERT INTO Products VALUES(4,'Corner Flags',2,34.95000000000000285);
COMMIT;

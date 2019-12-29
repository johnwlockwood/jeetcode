-- Stored procedure with help from http://www.mysqltutorial.org/mysql-stored-procedure-tutorial.aspx
DROP TABLE IF EXISTS Grades;

CREATE TABLE Grades (
    id INT(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `grade` INT(3) DEFAULT 0,
    PRIMARY KEY (`id`)
);

INSERT INTO Grades (`name`, `grade`) VALUES
("john", 50),
("jill", 87),
("mary", 80),
("kent", 55),
("pete", 60),
("john", 100),
("jill", 98),
("mary", 78),
("kent", 55),
("pete", 67),
("john", 99),
("jill", 98),
("mary", 76),
("kent", 55),
("pete", 80);

DROP PROCEDURE IF EXISTS GetTopHalfAvg;

DELIMITER $$

CREATE PROCEDURE GetTopHalfAvg()
BEGIN
    DECLARE gradeCount INT DEFAULT 0;
    DECLARE half INT DEFAULT 0;

    CREATE TEMPORARY TABLE WorkdPad (Name varchar(50) NOT NULL,
    GradeAvg INT) ENGINE=MEMORY; 

    INSERT INTO WorkdPad
    (Name, GradeAvg)
    SELECT `name`, AVG(`grade`) FROM Grades GROUP BY `name`;

    SELECT count(*) INTO gradeCount from WorkdPad;
    IF gradeCount > 0 THEN
        SET half = gradeCount/2;
    END IF;

    SELECT Name, GradeAvg FROM WorkdPad ORDER BY GradeAvg DESC LIMIT half;
    DROP TABLE WorkdPad;
END$$

DELIMITER ;

CALL  GetTopHalfAvg();

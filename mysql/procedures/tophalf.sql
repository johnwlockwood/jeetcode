-- Stored procedure with help from http://www.mysqltutorial.org/mysql-stored-procedure-tutorial.aspx
DROP TABLE IF EXISTS Grades;

CREATE TABLE Grades (
    id INT(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) DEFAULT NULL,
    `grade` INT(3) DEFAULT 0,
    PRIMARY KEY (`id`)
);

INSERT INTO Grades (`name`, `grade`) VALUES
("john", 100),
("jill", 98),
("mary", 80),
("Kent", 55),
("pete", 60);

DROP PROCEDURE IF EXISTS GetTopHalf;

DELIMITER $$

CREATE PROCEDURE GetTopHalf()
BEGIN
    DECLARE gradeCount INT DEFAULT 0;
    DECLARE half INT DEFAULT 0;

    SELECT count(*) INTO gradeCount from Grades;
    IF gradeCount > 0 THEN
        SET half = gradeCount/2;
    END IF;

    SELECT `name`, `grade` FROM Grades ORDER BY `grade` DESC LIMIT half;
END$$

DELIMITER ;

CALL  GetTopHalf();

CREATE TABLE table_stats (
	table_name varchar(64),
	cnt int(11),
	PRIMARY KEY (table_name)
);

REPLACE INTO table_stats VALUES('t_chatroom', (select count(*) from t_chatroom));

delimiter |
CREATE TRIGGER t_chatroom_stats_insert_trigger AFTER INSERT on t_chatroom for each row begin
  UPDATE table_stats SET cnt = cnt + 1 WHERE table_name = 't_chatroom';
END|

CREATE TRIGGER t_chatroom_stats_update_trigger AFTER DELETE on t_chatroom for each row begin
  UPDATE table_stats SET cnt = cnt - 1 WHERE table_name = 't_chatroom';
END|
delimiter ;
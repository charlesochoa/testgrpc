
CREATE TABLE IF NOT EXISTS user (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
)  ENGINE=INNODB;

insert  into `user`(`name`) values 

    ('Charles Ochoa'),
    ('Julia Pacheco');

CREATE TABLE IF NOT EXISTS test_media (
    test_media_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
)  ENGINE=INNODB;


insert  into `test_media`(`test_media_id`, `title`) values 

    (1, 'Hobbi Consolas'),
    (2, 'Top Gear');

CREATE TABLE IF NOT EXISTS test_gadget (
    test_gadget_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    test_media_id VARCHAR(255) NOT NULL,
    CONSTRAINT `media_ibfk_1` FOREIGN KEY (`test_media_id`) REFERENCES `test_media` (`test_media_id`)
)  ENGINE=INNODB;

insert  into `test_gadget`( `test_gadget_id`, `test_media_id`, `title`) values 

    (1, 1, 'Entre Parrafos'),
    (2, 1,'HC Detalle noticias'),
    (3, 2, 'Lecturas Recomendadas'),
    (4, 2,'Entre Parrafos');

CREATE TABLE IF NOT EXISTS test_click_impressions (
    test_click_id INT AUTO_INCREMENT PRIMARY KEY,
    test_gadget_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    CONSTRAINT `media_ibfk_2` FOREIGN KEY (`test_gadget_id`) REFERENCES `test_gadget` (`test_gadget_id`)
)  ENGINE=INNODB;


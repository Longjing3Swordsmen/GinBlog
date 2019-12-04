CREATE TABLE `post` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `create_time` TIMESTAMP default (datetime('now', 'localtime')),
    `update_time` TIMESTAMP default (datetime('now', 'localtime')),
    `title` VARCHAR(64) NOT NULL,
    `summary` VARCHAR(100) NULL,
    `content` TEXT NOT NULL,
    `category_id` INTEGER NOT NULL,
    FOREIGN KEY (category_id)
    REFERENCES category(id)
);

CREATE TABLE `category` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `create_time` TIMESTAMP default (datetime('now', 'localtime')),
    `update_time` TIMESTAMP default (datetime('now', 'localtime')),
    `name` VARCHAR(50) NOT NULL
);

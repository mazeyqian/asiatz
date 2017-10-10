<?php
ini_set("memory_limit", "100M");
require dirname(__FILE__).'/../core/init.php';

/* Do NOT delete this comment */
/* 不要删除这段注释 */

$configs = array(
    'name' => '三九养生堂',
    'tasknum' => 1,
    'domains' => array(
        'www.39yst.com',
        '39yst.com'
    ),
    'scan_urls' => array(
        'http://www.39yst.com'
    ),
    'list_url_regexes' => array(
        'http://www.39yst.com/\w+/',
        'http://www.39yst.com/\w+/\w+/'
    ),
    'content_url_regexes' => array(
        'http://www.39yst.com/\w+/.+\.shtml',
    )

);
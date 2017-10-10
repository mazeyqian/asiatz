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
    /* 'list_url_regexes' => array(
        'http://www.39yst.com/\w+/',
        'http://www.39yst.com/\w+/\w+/'
    ), */
    'content_url_regexes' => array(
        'http://www.39yst.com/\w+/.+\.shtml',
    ),
    'max_try' => 5,
    'export' => array(
        'type' => 'db',
        'table' => '39yst-1'
    ),
    'fields' => array(
        array(
            'name' => 'article_title',
            'selector_type' => 'regex',
            'selector' => '#<h1\s?>(.+)</h1>#i',
            'required' => true
        ),
        array(
            'name' => 'article_content',
            'selector_type' => 'regex',
            'selector' => '#\s?39yst\s?body\s?start\s?(.*)\s?39yst\s?body\s?end\s?#s',
            //'selector' => '#<h3>(.+)</h3>#i',
            /* 'selector' => "//div[@id='articleContent']", */
            'required' => true
        )
    )

);


$spider = new phpspider($configs);

$spider->start();
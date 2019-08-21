<?php

return [
    'class' => 'yii\db\Connection',
    'dsn' => 'pgsql:port=5432;host=127.0.0.1;dbname=planeplanner',
    'username' => 'planeplanner',
    'password' => 'planeplanner',
    'charset' => 'utf8',

    // Schema cache options (for production environment)
    //'enableSchemaCache' => true,
    //'schemaCacheDuration' => 60,
    //'schemaCache' => 'cache',
];

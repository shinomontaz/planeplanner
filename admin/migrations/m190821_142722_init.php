<?php

use yii\db\Migration;
use yii\db\Schema;

/**
 * Class m190821_142722_init
 */
class m190821_142722_init extends Migration
{
    // Use up()/down() to run migration code without a transaction.
    public function safeUp()
    {       
        $this->createTable('plane', [
            'id' => Schema::TYPE_PK,
            'name' => Schema::TYPE_STRING . ' NOT NULL',
            'cost' => Schema::TYPE_DECIMAL,
        ]);

        $this->createTable('brigade', [
            'id' => Schema::TYPE_PK,
            'name' => Schema::TYPE_STRING . ' NOT NULL',
        ]);

        $this->createTable('job', [
            'id' => Schema::TYPE_PK,
            'parent' => Schema::TYPE_INTEGER,
            'name' => Schema::TYPE_STRING . ' NOT NULL',
            'fk_brigade' => Schema::TYPE_INTEGER . ' NOT NULL',
            'time' => Schema::TYPE_INTEGER,
        ]);

        $this->createTable('plane_demand', [
            'fk_plane' => Schema::TYPE_INTEGER . ' NOT NULL',
            'fk_job' => Schema::TYPE_INTEGER . ' NOT NULL',
        ]);
        
        $this->createTable('schedule', [
            'fk_plane' => Schema::TYPE_INTEGER . ' NOT NULL',
            'arrival' => Schema::TYPE_TIMESTAMP . ' NOT NULL',
            'departure' => Schema::TYPE_TIMESTAMP,
        ]);
        
        $this->addForeignKey(
            'fk_plane_demand_plane',
            'plane_demand',
            'fk_plane',
            'plane',
            'id',
            'CASCADE'
        );

        $this->addForeignKey(
            'fk_plane_demand_job',
            'plane_demand',
            'fk_job',
            'job',
            'id',
            'CASCADE'
        );

        $this->addForeignKey(
            'fk_job_brigade',
            'job',
            'fk_brigade',
            'brigade',
            'id',
            'CASCADE'
        );

        $this->addForeignKey(
            'fk_job_parent',
            'job',
            'parent',
            'job',
            'id',
            'CASCADE'
        );

        $this->addForeignKey(
            'fk_schedule_plane',
            'schedule',
            'fk_plane',
            'plane',
            'id',
            'CASCADE'
        );
    }

    public function safeDown()
    {
        $this->dropForeignKey(
            'fk_plane_demand_plane',
            'plane_demand'
        );
        $this->dropForeignKey(
            'fk_job_brigade',
            'job'
        );
        $this->dropForeignKey(
            'fk_job_parent',
            'job'
        );
        $this->dropForeignKey(
            'fk_schedule_plane',
            'schedule'
        );
        $this->dropForeignKey(
            'fk_plane_demand_job',
            'plane_demand'
        );

        $this->dropTable(
            'job'
        );
        $this->dropTable(
            'brigade'
        );
        $this->dropTable(
            'schedule'
        );
        $this->dropTable(
            'plane'
        );
        $this->dropTable(
            'plane_demand'
        );
    }
}

<?php

namespace app\models;

use Yii;

/**
 * This is the model class for table "job".
 *
 * @property int $id
 * @property int $parent
 * @property string $name
 * @property int $fk_brigade
 * @property int $time
 *
 * @property Brigade $fkBrigade
 * @property Job $parent0
 * @property Job[] $jobs
 * @property PlaneDemand[] $planeDemands
 */
class Job extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'job';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['parent', 'fk_brigade', 'time'], 'default', 'value' => null],
            [['parent', 'fk_brigade', 'time'], 'integer'],
            [['name', 'fk_brigade'], 'required'],
            [['name'], 'string', 'max' => 255],
            [['fk_brigade'], 'exist', 'skipOnError' => true, 'targetClass' => Brigade::className(), 'targetAttribute' => ['fk_brigade' => 'id']],
            [['parent'], 'exist', 'skipOnError' => true, 'targetClass' => Job::className(), 'targetAttribute' => ['parent' => 'id']],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'parent' => 'Parent',
            'name' => 'Name',
            'fk_brigade' => 'Fk Brigade',
            'time' => 'Time',
        ];
    }

    /**
     * @return \yii\db\ActiveQuery
     */
    public function getFkBrigade()
    {
        return $this->hasOne(Brigade::className(), ['id' => 'fk_brigade']);
    }

    /**
     * @return \yii\db\ActiveQuery
     */
    public function getParent0()
    {
        return $this->hasOne(Job::className(), ['id' => 'parent']);
    }

    /**
     * @return \yii\db\ActiveQuery
     */
    public function getJobs()
    {
        return $this->hasMany(Job::className(), ['parent' => 'id']);
    }

    /**
     * @return \yii\db\ActiveQuery
     */
    public function getPlaneDemands()
    {
        return $this->hasMany(PlaneDemand::className(), ['fk_job' => 'id']);
    }
}

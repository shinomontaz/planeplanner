<?php

namespace app\models;

use Yii;

/**
 * This is the model class for table "plane".
 *
 * @property int $id
 * @property string $name
 * @property string $cost
 *
 * @property PlaneDemand[] $planeDemands
 * @property Schedule[] $schedules
 */
class Plane extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'plane';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['name'], 'required'],
            [['cost'], 'number'],
            [['name'], 'string', 'max' => 255],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'name' => 'Name',
            'cost' => 'Cost',
        ];
    }

    /**
     * @return \yii\db\ActiveQuery
     */
    public function getPlaneDemands()
    {
        return $this->hasMany(PlaneDemand::className(), ['fk_plane' => 'id']);
    }

    /**
     * @return \yii\db\ActiveQuery
     */
    public function getSchedules()
    {
        return $this->hasMany(Schedule::className(), ['fk_plane' => 'id']);
    }
}

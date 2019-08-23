<?php

namespace app\models;

use Yii;

/**
 * This is the model class for table "schedule".
 *
 * @property int $fk_plane
 * @property string $arrival
 * @property string $departure
 *
 * @property Plane $fkPlane
 */
class Schedule extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'schedule';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['fk_plane', 'arrival'], 'required'],
            [['fk_plane'], 'default', 'value' => null],
            [['fk_plane'], 'integer'],
            [['arrival', 'departure'], 'safe'],
            [['fk_plane'], 'exist', 'skipOnError' => true, 'targetClass' => Plane::className(), 'targetAttribute' => ['fk_plane' => 'id']],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'fk_plane' => 'Fk Plane',
            'arrival' => 'Arrival',
            'departure' => 'Departure',
        ];
    }

    /**
     * @return \yii\db\ActiveQuery
     */
    public function getFkPlane()
    {
        return $this->hasOne(Plane::className(), ['id' => 'fk_plane']);
    }
}

<?php

use yii\helpers\Html;

/* @var $this yii\web\View */
/* @var $model app\models\Brigade */

$this->title = 'Create Brigade';
$this->params['breadcrumbs'][] = ['label' => 'Brigades', 'url' => ['index']];
$this->params['breadcrumbs'][] = $this->title;
?>
<div class="brigade-create">

    <h1><?= Html::encode($this->title) ?></h1>

    <?= $this->render('_form', [
        'model' => $model,
    ]) ?>

</div>

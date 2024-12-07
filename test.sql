select * from (
select --top 100
  --COALESCE(tf.doc_fix_date, tt.doc_date) as 'history_date',
  tf.doc_fix_date as 'history_date',
  'ТТН' as 'history_type',
  IIF(tt.doc_type='Возврат от меня','Возврат покупателю', tt.ttn_type) as 'history_reason',
  tt.id as 'document_id',
  tp.id as 'document_position_id',
  tp.product_identity AS 'document_position_identity',
  tt.doc_number as 'document_number',
  tt.doc_date as 'document_date',
  tt.status as 'document_status',
  coalesce((SELECT concat(ttn_number,' ',ttn_date) FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id),concat('отсутствует ',tp.product_inform_f2_reg_id)) as 'document_source',
  tp.product_unit_type as 'document_position_unit_type',
  COALESCE(cast(tp.product_quantity as float),0) as 'document_position_quantity',
  COALESCE(cast(tp.product_quantity as float),0)  as 'document_position_quantity_fact',
  tp.product_full_name as 'document_position_full_name',
  tp.product_code as 'document_position_code',
  tp.product_alc_code as 'document_position_alc_code',
  cast(tp.product_alc_volume as float) as 'document_position_alc_volume',
  coalesce((SELECT CAST(product_alc_volume AS float) FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), 0.0) as 'document_position_alc_volume_fact',
  cast(tp.product_capacity as float) as 'document_position_capacity',
  tt.consignee_client_reg_id as 'document_partner_fsrar_id',
  tp.producer_client_reg_id as 'document_position_producer_fsrar_id',
  coalesce((SELECT product_inform_f1_reg_id FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'document_position_form1',
  coalesce((SELECT product_inform_f2_reg_id FROM form2_egais WHERE product_inform_f2_reg_id = tp.product_inform_f2_reg_id ), concat('отсутствует ', tp.product_inform_f2_reg_id)) as 'document_position_form2',
  COALESCE(tf.doc_fix_number, '') as 'document_fix_number',
  COALESCE(tf.doc_reg_id, '') as 'document_reg_id',
  coalesce((SELECT bottling_date FROM form1_egais WHERE product_inform_f1_reg_id = tp.product_inform_f1_reg_id ), concat('отсутствует ', tp.product_inform_f1_reg_id)) as 'document_position_bottling_date',
  cast(case tp.product_unit_type when 'Packed' then COALESCE(cast(tp.product_quantity as float),0) else 0 end as int) as 'document_position_counts',
  case tp.product_unit_type when 'Unpacked' then COALESCE(cast(tp.product_quantity as float),0) else 0 end as 'document_position_dal',
  round(case tp.product_unit_type when 'Unpacked' then COALESCE(cast(tp.product_quantity as float),0) 
        else  COALESCE(cast(tp.product_quantity as float),0) * cast(tp.product_capacity as float) * 0.1 end,4) as 'document_position_volume',
  COALESCE(cast(tp.product_quantity as float),0) * cast(tp.product_price as float) as 'document_position_summ'
FROM ttn tt 
  join ttn_products tp on tp.id_ttn = tt.id
  join ttn_form2 tf on tf.id_ttn = tt.id
where 
  tt.ttn_type in ('Исходящий', 'Импорт')
ORDER by tf.doc_fix_date
) out 
where 
  history_date like '%2023.08.01%'
  and history_type like '%%'
LIMIT 100 OFFSET 100
--OFFSET 100 ROWS
--FETCH NEXT 100 ROWS ONLY;
;
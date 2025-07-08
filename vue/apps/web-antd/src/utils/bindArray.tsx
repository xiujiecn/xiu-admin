/**
 * 编辑form 中的数组元素
 * @param editFormApi  formApi 
 * @param key  数组key
 * @param oldData  旧数组数据
 * @param index  索引
 * @param itemData  index 位置的元素的新数据
 */
export async function editArray( editFormApi:any, key:string, oldData:any, index:number, itemData:any) {
    let newData = oldData.map((item:any, i:number) => {
      if (i === index) {
        return itemData;
      }
      return item;
    });
    await editFormApi.setFieldValue(key, newData);
  }
  
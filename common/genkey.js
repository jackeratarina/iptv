import { Md5 } from "https://deno.land/std@0.119.0/hash/md5.ts";
import { randomNumber } from "https://deno.land/x/random_number/mod.ts";



function MD5(str){
    const md5 = new Md5();
    return md5.update(str).toString();
    }
function getHeadersAuthen(){
    let currentTimestamp = Date.now()
    let timeDiff = - randomNumber({ min: 99, max: 199 });

  
    currentTimestamp = currentTimestamp - timeDiff
  
    const date = new Date(currentTimestamp)
    const day = date.getDate() < 10 ? ('0' + date.getDate()) : (date.getDate()).toString()
    const month = (date.getMonth() + 1) < 10 ? ('0' + (date.getMonth() + 1)) : (date.getMonth() + 1).toString()
    const year = date.getFullYear().toString()
    const hour = date.getHours() < 10 ? ('0' + date.getHours()) : (date.getHours()).toString()
    const minute = date.getMinutes() < 10 ? ('0' + date.getMinutes()) : (date.getMinutes()).toString()
    const second = date.getSeconds() < 10 ? ('0' + date.getSeconds()) : (date.getSeconds()).toString()
  
    const dateValue = year + month + day
    const timeValue = hour + minute + second
    const md5Value = (MD5(dateValue + timeValue)).toString()
    const keyValue = md5Value.substring(0, 3) + md5Value.substring(md5Value.length - 3)
  
    const keyAccess = "Kh0ngDuLieu" + dateValue + "C0R0i" + timeValue + "Kh0aAnT0an" + keyValue
  
  return {
    'X-SFD-Key': MD5(keyAccess).toString(),
    'X-SFD-Date': dateValue + timeValue,
          "Content-Type": "application/json",
          'User-Agent' : 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36',
  };
}
async function getData(channel){
    // https://api.thvli.vn/backend/cm/get_detail/thvl1-hd/?timezone=Asia/Saigon
    const resp = await fetch(`https://api.thvli.vn/backend/cm/get_detail/${channel}-hd/?timezone=Asia/Saigon`, {
        method: "GET",
        headers: getHeadersAuthen()
      });
      return resp;
}
export default getData
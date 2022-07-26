let BASE_URL = 'http://app:8080/newsletter/file';
  export async function StoreFile(newsletter) {
    const response = await fetch(BASE_URL, {
        method: 'POST',        
        body: newsletter,
        redirect: 'follow'
      })
    return await response.json();
}
  
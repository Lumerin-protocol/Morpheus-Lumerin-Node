export const getSessionsByUser = async (url, user) => {
    if(!user || !url) {
      return;
    }

    const getSessions = async (user, offset, limit) => {
      try {
        const path = `${url}/blockchain/sessions/user?user=${user}&offset=${offset}&limit=${limit}&order=desc`;
        const response = await fetch(path);
        const data = await response.json();
        return data.sessions;
      }
      catch (e) {
        console.log("Error", e)
        return [];
      }
    } 
    
    const limit = 20;
    let offset = 0;
    let sessions: any[] = [];
    let all = false;

    while (!all) {
      console.log("Getting session for user: ", user, offset, limit)
      const sessionsRes = await getSessions(user, offset, limit);
      sessions.push(...sessionsRes);

      if(sessionsRes.length != limit) {
        all = true;
      }
      else {
        offset++;
      }
    }

    return sessions;
}

export const getBidsByModelId = async (url, modelId) => {
  if(!modelId || !url) {
    return;
  }

  const getBidsByModels = async (modelId, offset, limit) => {
    try {
      const path = `${url}/blockchain/models/${modelId}/bids?offset=${offset}&limit=${limit}&order=desc`
      const response = await fetch(path);
      const data = await response.json();
      return data.bids;
    }
    catch (e) {
      console.log("Error", e)
      return [];
    }
  }
  
  const limit = 20;
  let offset = 0;
  let bids: any[] = [];
  let all = false;

  while (!all) {
    console.log("Getting bids by model id: ", modelId, offset, limit)
    const bidsRes = await getBidsByModels(modelId, offset, limit);
    bids.push(...bidsRes);

    if(bids.length != limit) {
      all = true;
    }
    else {
      offset++;
    }
  }

  return bids;
}

export const getBidInfoById = async (url, id) => {
  try {
    const path = `${url}/blockchain/bids/${id}`
    const response = await fetch(path);
    const data = await response.json();
    return data.bid;
  }
  catch (e) {
    console.log("Error", e)
    return undefined;
  }
}
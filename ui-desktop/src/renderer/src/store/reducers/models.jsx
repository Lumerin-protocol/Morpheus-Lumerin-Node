import { handleActions } from 'redux-actions';

const initialState = {
    selectedModel: null,
    selectedBid: null,
    selectedProvider: null,
    activeSession: null
};

const reducer = handleActions(
    {
        'set-model': (state, { payload }) => ({
            ...state,
            selectedModel: payload
        }),
        'set-bid': (state, { payload }) => ({
            ...state,
            selectedBid: payload.bidId,
            selectedProvider: payload.provider
        }),
        'set-active-session': (state, { payload }) => ({
            ...state,
            activeSession: payload,
        }),
    },
    initialState
);

export default reducer;

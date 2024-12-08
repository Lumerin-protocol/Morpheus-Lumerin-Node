import { useState } from 'react';
import styled from 'styled-components';
import Select from "react-select";
import {
    RightBtn,
} from '../../contracts/modals/CreateContractModal.styles';
import { abbreviateAddress } from '../../../utils';
import { formatSmallNumber } from '../utils';
import { IconX, IconPlugConnectedX } from '@tabler/icons-react';

const RowContainer = styled.div`
  padding: 0 1.2rem;
  height: 40px;
  display: grid;
  grid-template-columns: 3fr 1fr 160px;
  text-align: center;
  border: ${p => p.theme.colors.morMain} solid 0.5px;
  color: ${p => p.theme.colors.morMain};
  background: rgba(0,0,0, 0.1);
  border-radius: 5px;
  margin-bottom: 5px;

  &:last-child {
    margin-bottom: 0
  }
`;

const FlexCenter = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
`;

const Buttons = styled.div`
    display: flex;
    justify-content: end;
    min-width: 150px;
    width: 150px;
    align-items: center;
    gap: 10px;

    button {
        height: auto;
    }
`

const PriceContainer = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    white-space: nowrap;
`

const ModelNameContainer = styled(FlexCenter)`
    justify-content: start;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
`

const DisconnectedIcon = styled(IconPlugConnectedX)`
    width: 16px;
    color: white;
    margin-left: 10px;
`

const selectorStyles = {
    control: (base) => ({ ...base, borderColor: '#20dc8e', width: '100%', background: 'transparent' }),
    option: (base, state) => ({
        ...base,
        backgroundColor: state.isSelected ? '#0e4353' : "#03160e",
        color: '#20dc8e',
        ':active': {
            ...base[':active'],
            backgroundColor: '#0e435380',
            color: '#20dc8e'
        },
        zIndex: 100
    }),
    placeholder: base => ({
        ...base,
        color: '#20dc8e',
        fontSize: '1.4rem',
        fontWeight: 700
    }),
    singleValue: base => ({
        ...base,
        color: '#20dc8e',
        fontWeight: 600,
        fontSize: '1.4rem'
    }),
    indicatorsContainer: base => ({
        ...base,
        color: '#20dc8e'
    }),
};

function ModelRow(props) {
    const bids = props?.model?.bids || [];
    const modelId = props?.model?.Id || '';
    const isLocal = props?.model?.isLocal;
    const lastAvailabilityCheck: Date = new Date(props?.model?.lastCheck ?? new Date());

    const [selected, changeSelected] = useState<any>();
    const [useSelect, setUseSelect] = useState<boolean>();
    const [targetBid, setTargetBid] = useState<any>();

    const options = bids.map(x => {
        return ({ value: x.Id, label: `${abbreviateAddress(x.Provider || "", 3)} ${formatSmallNumber(x.PricePerSecond / (10 ** 18))} ${props.symbol}` })
    });

    const handleChangeModel = () => {
        props.onChangeModel({ modelId, bidId: targetBid?.Id })
    }

    const selectLocal = () => {
        props.onChangeModel({ modelId, isLocal: true })
    }

    const onChangeSelector = (data) => {
        const bid = bids.find(x => x.Id == data.value);
        setTargetBid({ ...bid, customSelection: true });
    }

    const formatPrice = () => {
        if (targetBid) {
            return `${formatSmallNumber(targetBid?.PricePerSecond / (10 ** 18))} ${props.symbol}`;
        }

        const prices = bids.filter(x => x.Id).map(x => Number(x.PricePerSecond));
        if (prices.length == 1) {
            return `${formatSmallNumber(prices[0] / (10 ** 18))} ${props.symbol}`;
        }

        const minPrice = Math.min(...prices);
        const maxPrice = Math.max(...prices);

        return `${formatSmallNumber(minPrice / (10 ** 18))} - ${formatSmallNumber(maxPrice / (10 ** 18))} ${props.symbol}`
    }

    return (
        <RowContainer useSelect={useSelect}>
            <ModelNameContainer>
                { props?.model?.Name } 
                { 
                    !props?.model?.isOnline && 
                    <DisconnectedIcon data-rh-negative data-rh={`Last seen offline at ${lastAvailabilityCheck?.toLocaleTimeString()}`} /> 
                }
            </ModelNameContainer>
            <PriceContainer>
                {
                    useSelect
                        ? <div style={{ width: '100%', display: 'flex', alignItems: 'center' }}>
                            <div style={{ marginRight: '10px', width: '95%' }}>
                                <Select
                                    menuPlacement='auto'
                                    onChange={onChangeSelector}
                                    styles={selectorStyles}
                                    value={selected}
                                    options={options}></Select>
                            </div>

                            <IconX width={'1.5rem'} style={{ cursor: 'pointer' }} onClick={() => {
                                setUseSelect(!useSelect);
                                setTargetBid(undefined);
                                changeSelected(undefined);
                            }}></IconX>
                        </div>
                        : <div>
                            {
                                !isLocal ? (
                                    <FlexCenter>
                                        <span
                                            data-rh-negative
                                            data-rh={`Bid ID: ${abbreviateAddress(targetBid?.Id || "", 5)}`}
                                            style={{ marginRight: '10px' }}>
                                            {formatPrice()}
                                        </span>
                                        {/* <IconEdit width={'1.5rem'} style={{ cursor: 'pointer' }} onClick={() => setUseSelect(!useSelect)}></IconEdit> */}
                                    </FlexCenter>
                                    ) : 
                                    <FlexCenter>(local)</FlexCenter>
                            }
                        </div>
                }
            </PriceContainer>
            <Buttons>
                <RightBtn block disabled={!props?.model?.isOnline && !isLocal} onClick={() => isLocal ? selectLocal() : handleChangeModel()}>Select</RightBtn>
            </Buttons>
        </RowContainer>
    );
}

export default ModelRow;

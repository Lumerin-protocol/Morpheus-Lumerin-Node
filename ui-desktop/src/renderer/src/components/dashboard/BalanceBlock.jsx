import React, { useState, useContext } from 'react';
import withBalanceBlockState from '../../store/hocs/withBalanceBlockState';
import { EtherIcon } from '../icons/EtherIcon';
import { LumerinLogoFull } from '../icons/LumerinLogoFull';
import { Balance } from './Balance';
import {
  WalletBalanceHeader,
  Btn,
  BtnAccent,
  BtnRow,
  SecondaryContainer,
  Container,
  Primary,
  CoinsRow,
  BalanceContainer,
  GlobalContainer
} from './BalanceBlock.styles';
import Spinner from '../common/Spinner';
import { ToastsContext } from '../toasts';

const WalletBalance = ({
  lmrBalance,
  lmrBalanceUSD,
  ethBalance,
  ethBalanceUSD,
  symbol,
  symbolEth
}) => (
  <BalanceContainer>
    <CoinsRow>
      <Primary data-testid="lmr-balance">
        <Balance
          currency={symbol}
          value={lmrBalance}
          icon={
            <LumerinLogoFull style={{ color: 'white', height: "2rem"}}/> 
          }
          equivalentUSD={lmrBalanceUSD}
          maxSignificantFractionDigits={0}
        />
      </Primary>
      <Primary data-testid="eth-balance">
        <Balance
          currency={symbolEth}
          value={ethBalance}
          icon={<EtherIcon size="3.3rem" />}
          equivalentUSD={ethBalanceUSD}
          maxSignificantFractionDigits={5}
        />
      </Primary>
    </CoinsRow>
  </BalanceContainer>
);

const BalanceBlock = ({
  lmrBalance,
  lmrBalanceUSD,
  ethBalance,
  ethBalanceUSD,
  onTabSwitch,
  symbol,
  symbolEth,
  ...props
}) => {
  const handleTabSwitch = e => {
    e.preventDefault();
    onTabSwitch(e.target.dataset.modal);
  };

  return (
    <GlobalContainer>
      <Container>
        <SecondaryContainer>
          <WalletBalance
            {...{
              lmrBalance: props?.balances?.mor ? +props.balances.mor / 10 ** 18 : 0,
              lmrBalanceUSD: props?.balances?.mor ? `$${((+props.balances.mor / 10 ** 18) * +props.rate).toFixed(0)}` : 0,
              ethBalance:  props?.balances?.eth ? (+props.balances.eth / 10 ** 18) : 0,
              ethBalanceUSD,
              symbol,
              symbolEth
            }}
          />
          <BtnRow>
            <BtnAccent
              data-modal="receive"
              data-testid="receive-btn"
              onClick={handleTabSwitch}
              block
            >
              Receive
            </BtnAccent>
            <BtnAccent
              data-modal="send"
              data-testid="send-btn"
              onClick={handleTabSwitch}
              block
            >
              Send
            </BtnAccent>
          </BtnRow>
        </SecondaryContainer>
      </Container>
    </GlobalContainer>
  );
};

export default withBalanceBlockState(BalanceBlock);

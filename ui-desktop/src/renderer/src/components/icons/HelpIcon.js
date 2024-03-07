import React from 'react';
import styled from 'styled-components';

import BaseIcon from './BaseIcon';

const HelpIcon = ({ isActive, size }, props) => {
  const Circle = styled.div`
    background-color: ${p => p.theme.colors.medium};
    border-radius: 50%;
  `;
  const fill = isActive ? '#11B4BF' : 'black';

  return (
    <Circle>
      <BaseIcon size={size} viewBox="0 0 40 40" {...props}>
        <path
          d="M19.9997 12C17.2431 12 15 14.2431 15 16.9997C15 17.7875 15.6407 18.4282 16.4286 18.4282C17.2164 18.4282 17.8571 17.7875 17.8571 16.9997C17.8571 15.8182 18.8182 14.8571 19.9997 14.8571C21.1818 14.8571 22.1422 15.8182 22.1422 16.9997C22.1422 17.9641 21.4946 18.8143 20.5651 19.0676C19.3911 19.3879 18.5711 20.3976 18.5711 21.5236V23.1904C18.5711 23.9782 19.2118 24.6196 19.9997 24.6196C20.7868 24.6196 21.4282 23.9782 21.4282 23.1904V21.7933C23.5372 21.1649 25 19.212 25 17.001C25 14.2431 22.7562 12 19.9997 12Z"
          fill={fill}
        />
        <path
          d="M19.9998 28.4282C20.7888 28.4282 21.4284 27.7886 21.4284 26.9996C21.4284 26.2106 20.7888 25.571 19.9998 25.571C19.2108 25.571 18.5712 26.2106 18.5712 26.9996C18.5712 27.7886 19.2108 28.4282 19.9998 28.4282Z"
          fill={fill}
        />
      </BaseIcon>
    </Circle>
  );
};

export default HelpIcon;

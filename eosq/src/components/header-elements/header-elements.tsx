import * as React from "react"
import { color as color_, fontSize } from "styled-system"
import { Link } from "react-router-dom"
import { Cell, Grid } from "../../atoms/ui-grid/ui-grid.component"
import { Links } from "../../routes"
import { t } from "i18next"
import { styled } from "../../theme"
import { Config } from "../../models/config"
import { Box, Text } from "@dfuse/explorer"
import { Img } from "../../atoms/img"

const LogoElement: React.ComponentType<any> = styled.div`
  font-family: "Lato", sans-serif;
  font-weight: 600;
  ${color_};
  ${fontSize};
  top: -10px;
  position: relative;

  @media (max-width: 767px) {
    top: -6px;
  }
`

const Tagline: React.ComponentType<any> = styled.span`
  font-family: "Lato", sans-serif;
  font-weight: 600;
  color: ${(props) => props.theme.colors.logo1};
  ${fontSize};
  letter-spacing: 1px;
`

const LogoLink: React.ComponentType<any> = styled(Link)`
  display: block;
  display: flex;
  align-items: center;
  justify-content: center;
`

interface Props {
  variant: "dark" | "light"
}

export const HeaderLogo: React.FC<Props> = () => {
  return (
    <Grid gridTemplateColumns={["auto auto"]} gridRow={["1"]} gridColumn={["1"]} py={[1, 0]}>
      <Cell alignSelf="center" justifySelf="right">
        <LogoLink to={Links.home()}>
          <Logo />
        </LogoLink>
      </Cell>
    </Grid>
  )
}

const Logo: React.FC = () => {
  const { network } = Config

  return (<Img src={"/images/logo-explorer-black.png"} alt="Logo" minWidth="70px" maxHeight="70px" maxWidth="70vw"></Img>)
}

const LogoDefault: React.FC = () => (
  <>
a  </>
)

const LogoImage: React.FC<{ image: string }> = ({ image }) => (
  <Img src={image} alt="Logo" minWidth="70px" maxHeight="70px" maxWidth="70vw"></Img>
)

const LogoText = styled(Text)`
  font-family: "Lato", sans-serif;
  font-weight: 400;
`

const LogoImageAndText: React.FC<{ image: string; text?: string }> = ({ image, text }) => (
  <Box pa={[0]} alignItems="center" justifyContent="center" minWidth="150px" flexWrap="wrap">
    <Img src={image} alt="Logo" title={text} maxWidth="48px" maxHeight="48px"></Img>
    {text ? (
      text === "eosq" ? (
        <Box mx={[2]}>
          <LogoDefault />
        </Box>
      ) : (
        <LogoText color="white" mx={[2]} fontSize={[4]}>
          {text}
        </LogoText>
      )
    ) : null}
  </Box>
)

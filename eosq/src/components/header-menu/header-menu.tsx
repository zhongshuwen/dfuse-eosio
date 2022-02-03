import * as React from "react"
import { Cell, Grid } from "../../atoms/ui-grid/ui-grid.component"
import { Text } from "../../atoms/text/text.component"
import { HeaderLogo } from "../header-elements/header-elements"
import { MainMenu } from "../main-menu/main-menu.component"
import { theme, styled } from "../../theme"
import { NetworkSelector } from "../settings-selectors/network-selector"
import { LanguageSelector } from "../settings-selectors/language-selector"
import { t } from "i18next"
import { getCurrentLanguageValue } from "../settings-selectors/settings.helpers"
import { Img } from "../../atoms/img"

const HeaderWrapper: React.ComponentType<any> = styled(Cell)`
  width: 100%;



  background: #111; 
  background: linear-gradient(
    to right,
    #000000 8%,
    #1a1a1d 93%
  );
`

export class HeaderMenu extends React.Component {
  renderSectionTitle(text: string) {
    return (
      <Text pb={[2]} pl={[3]} fontWeight="600" fontSize={[3]} color={"#fff"}>
        {text}
      </Text>
    )
  }

  render() {
    return (
      <HeaderWrapper mx="auto">
        <Cell mx="auto" px={[2, 3, 4]} py={[0, 3]}>
          <Grid
            gridTemplateColumns={["1fr", "1fr 1fr 1fr 1fr", "1fr 1fr 1fr 1fr"]}
            pt={[0]}
            pb={[0]}
            px={[1, 0]}
            gridColumnGap={[0, 1, 2]}
          >
            <div style={{height:"100%", display:"flex", justifyContent:"center",alignItems:"center"}} > 
            <Img src={"/images/logo-explorer-white.png"} alt="Logo" minWidth="70px" maxHeight="70px" maxWidth="70vw"></Img>
             
            </div>

            <Cell
              height="100%"
              borderLeft={`2px solid ${theme.colors.bleu6}`}
              alignSelf="right"
              justifySelf={["inline", "inline"]}
              px={[0, 4]}
              py={[2]}
            >
              {this.renderSectionTitle(t("core.menu.titles.navigation"))}
              <MainMenu variant="light" />
            </Cell>
            <Cell
              height="100%"
              borderLeft={`2px solid ${theme.colors.bleu6}`}
              alignSelf="right"
              justifySelf={["inline", "inline"]}
              px={[0, 4]}
              py={[2]}
            >
              {this.renderSectionTitle(t("core.menu.titles.network"))}
              <NetworkSelector variant="light" />
            </Cell>
            <Cell
              height="100%"
              borderLeft={`2px solid ${theme.colors.bleu6}`}
              alignSelf="right"
              justifySelf={["inline", "inline"]}
              px={[0, 4]}
              py={[2]}
            >
              {this.renderSectionTitle(t("core.menu.titles.language"))}
              <LanguageSelector variant="light" />
            </Cell>
          </Grid>
        </Cell>
      </HeaderWrapper>
    )
  }
}

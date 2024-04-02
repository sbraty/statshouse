// This file was generated by lezer-generator. You probably shouldn't edit it.
import { LRParser } from '@lezer/lr';
import { extendIdentifier, specializeIdentifier } from './tokens';
import { promQLHighLight, promQLIndent } from './highlight';

const spec_Identifier = {
  __proto__: null,
  absent_over_time: 329,
  absent: 331,
  abs: 333,
  acos: 335,
  acosh: 337,
  asin: 339,
  asinh: 341,
  atan: 343,
  atanh: 345,
  avg_over_time: 347,
  ceil: 349,
  changes: 351,
  clamp: 353,
  clamp_max: 355,
  clamp_min: 357,
  cos: 359,
  cosh: 361,
  count_over_time: 363,
  days_in_month: 365,
  day_of_month: 367,
  day_of_week: 369,
  day_of_year: 371,
  deg: 373,
  delta: 375,
  deriv: 377,
  exp: 379,
  floor: 381,
  histogram_count: 383,
  histogram_fraction: 385,
  histogram_quantile: 387,
  histogram_stddev: 389,
  histogram_stdvar: 391,
  histogram_sum: 393,
  histogram_avg: 395,
  holt_winters: 397,
  hour: 399,
  idelta: 401,
  increase: 403,
  irate: 405,
  label_replace: 407,
  label_join: 409,
  last_over_time: 411,
  ln: 413,
  log10: 415,
  log2: 417,
  mad_over_time: 419,
  max_over_time: 421,
  min_over_time: 423,
  minute: 425,
  month: 427,
  pi: 429,
  predict_linear: 431,
  present_over_time: 433,
  quantile_over_time: 435,
  rad: 437,
  rate: 439,
  resets: 441,
  round: 443,
  scalar: 445,
  sgn: 447,
  sin: 449,
  sinh: 451,
  sort: 453,
  sort_desc: 455,
  sort_by_label: 457,
  sort_by_label_desc: 459,
  sqrt: 461,
  stddev_over_time: 463,
  stdvar_over_time: 465,
  sum_over_time: 467,
  tan: 469,
  tanh: 471,
  timestamp: 473,
  time: 475,
  vector: 477,
  year: 479,
};
export const parser = LRParser.deserialize({
  version: 14,
  states:
    "8fOYQPOOO'nQPOOOOQO'#Cz'#CzO'sQPO'#CyQ(OQQOOOOQO'#Da'#DaO'yQPO'#D`OOQO'#FT'#FTO)]QPO'#FZOYQPO'#FVOYQPO'#FYO/zQSO'#F]O0SQQO'#F[OOQO'#F['#F[OOQO'#Fq'#FqOOQO'#Db'#DbOOQO'#Dd'#DdOOQO'#De'#DeOOQO'#Df'#DfOOQO'#Dg'#DgOOQO'#Dh'#DhOOQO'#Di'#DiOOQO'#Dj'#DjOOQO'#Dk'#DkOOQO'#Dl'#DlOOQO'#Dm'#DmOOQO'#Dn'#DnOOQO'#Do'#DoOOQO'#Dp'#DpOOQO'#Dq'#DqOOQO'#Dr'#DrOOQO'#Ds'#DsOOQO'#Dt'#DtOOQO'#Du'#DuOOQO'#Dv'#DvOOQO'#Dw'#DwOOQO'#Dx'#DxOOQO'#Dy'#DyOOQO'#Dz'#DzOOQO'#D{'#D{OOQO'#D|'#D|OOQO'#D}'#D}OOQO'#EO'#EOOOQO'#EP'#EPOOQO'#EQ'#EQOOQO'#ER'#EROOQO'#ES'#ESOOQO'#ET'#ETOOQO'#EU'#EUOOQO'#EV'#EVOOQO'#EW'#EWOOQO'#EX'#EXOOQO'#EY'#EYOOQO'#EZ'#EZOOQO'#E['#E[OOQO'#E]'#E]OOQO'#E^'#E^OOQO'#E_'#E_OOQO'#E`'#E`OOQO'#Ea'#EaOOQO'#Eb'#EbOOQO'#Ec'#EcOOQO'#Ed'#EdOOQO'#Ee'#EeOOQO'#Ef'#EfOOQO'#Eg'#EgOOQO'#Eh'#EhOOQO'#Ei'#EiOOQO'#Ej'#EjOOQO'#Ek'#EkOOQO'#El'#ElOOQO'#Em'#EmOOQO'#En'#EnOOQO'#Eo'#EoOOQO'#Ep'#EpOOQO'#Eq'#EqOOQO'#Er'#ErOOQO'#Es'#EsOOQO'#Et'#EtOOQO'#Eu'#EuOOQO'#Ev'#EvOOQO'#Ew'#EwOOQO'#Ex'#ExOOQO'#Ey'#EyOOQO'#Ez'#EzOOQO'#E{'#E{OOQO'#E|'#E|OOQO'#E}'#E}OOQO'#FO'#FOOOQO'#FP'#FPOOQO'#FQ'#FQQOQPOOO1mQPO'#C{O1rQPO'#DOO'yQPO,59eO1yQQO,59eO3gQPO,59kO3gQPO,59kO3gQPO,59kO3gQPO,59kO3gQPO,59kO:UQQO,5;mO:ZQQO,5;pO:cQPO,5<POOQO,59z,59zOOQO,5;o,5;oO:zQQO,5;qO;RQQO,5;tO<iQSO'#F^OOQO,5;w,5;wO<zQPO,5;wOOQO,5;v,5;vO=SQSO'#C|OOQO,59g,59gOOQO,59j,59jO=[QQO,59jOOQO1G/P1G/POOQO'#DR'#DRO1mQPO'#DSOOQO'#Fu'#FuO=fQPO'#FuOYQPO1G/VOYQPO1G/VOYQPO1G/VOYQPO1G/VOYQPO1G/VODQQWO1G1XOOQO1G1[1G1[ODYQQO1G1[OD_QPO'#FTOOQO'#Fg'#FgOOQO1G1k1G1kODjQPO1G1kOOQO1G1]1G1]OOQO'#F_'#F_ODoQPO,5;xODwQSO1G1cOEPQPO1G1cOOQO1G1c1G1cOOQO,59h,59hOEXQPO,59hOYQPO'#FjOEaQPO1G/UOOQO1G/U1G/UOEiQPO,59nOOQO,5<a,5<aOM_QQO7+$qOMoQQO7+$qO! TQQO7+$qO! kQQO7+$qO!#SQQO7+$qOOQO7+&s7+&sO!#mQQO7+&yOOQO7+&v7+&vO!#uQPO7+'VOOQO1G1d1G1dOOQO,5<V,5<VOOQO7+&}7+&}O!#zQSO7+&}OOQO-E9i-E9iO!$SQSO1G/SO!$[QPO1G/SOOQO1G/S1G/SO!$dQQO,5<UOOQO-E9h-E9hOOQO7+$p7+$pO!$nQPO1G/YOOQO<<Je<<JeO!+SQPO<<JeOOQO<<Jq<<JqOOQO<<Ji<<JiP!+XQSO'#FkOOQO,5<T,5<TOOQO7+$n7+$nO!+^QSO7+$nOOQO-E9g-E9gOOQO7+$t7+$tOOQOAN@PAN@POOQO<<HY<<HYP!+fQSO'#Fi",
  stateData:
    '!+k~O$dOSkOS~OWQOXQOYQOZQO[QO]QO^QO_QO`QOaQObQOcQO!V[O#z^O$W^O$aVO$bVO$fXO$j_O$k`O$laO$mbO$ncO$odO$peO$qfO$rgO$shO$tiO$ujO$vkO$wlO$xmO$ynO$zoO${pO$|qO$}rO%OsO%PtO%QuO%RvO%SwO%TxO%UyO%VzO%W{O%X|O%Y}O%Z!OO%[!PO%]!QO%^!RO%_!SO%`!TO%a!UO%b!VO%c!WO%d!XO%e!YO%f!ZO%g![O%h!]O%i!^O%j!_O%k!`O%l!aO%m!bO%n!cO%o!dO%p!eO%q!fO%r!gO%s!hO%t!iO%u!jO%v!kO%w!lO%x!mO%y!nO%z!oO%{!pO%|!qO%}!rO&O!sO&P!tO&Q!uO&R!vO&S!wO&T!xO&U!yO&V!zO&W!{O&X!|O&[WO&]WO&^VO&`ZO~O!V!}O~Od#OOe#OO$f#PO~OU#YOV#SOf#VOg#WOh#VOt#SOw#SOx#SOy#SOz#TO{#TO|#UO}#UO!O#UO!P#UO!Q#UO!R#UO$Y#ZO&Y#XO~O$a#]O$b#]O&^#]OW#}XX#}XY#}XZ#}X[#}X]#}X^#}X_#}X`#}Xa#}Xb#}Xc#}X!V#}X#z#}X$W#}X$a#}X$b#}X$f#}X$j#}X$k#}X$l#}X$m#}X$n#}X$o#}X$p#}X$q#}X$r#}X$s#}X$t#}X$u#}X$v#}X$w#}X$x#}X$y#}X$z#}X${#}X$|#}X$}#}X%O#}X%P#}X%Q#}X%R#}X%S#}X%T#}X%U#}X%V#}X%W#}X%X#}X%Y#}X%Z#}X%[#}X%]#}X%^#}X%_#}X%`#}X%a#}X%b#}X%c#}X%d#}X%e#}X%f#}X%g#}X%h#}X%i#}X%j#}X%k#}X%l#}X%m#}X%n#}X%o#}X%p#}X%q#}X%r#}X%s#}X%t#}X%u#}X%v#}X%w#}X%x#}X%y#}X%z#}X%{#}X%|#}X%}#}X&O#}X&P#}X&Q#}X&R#}X&S#}X&T#}X&U#}X&V#}X&W#}X&X#}X&[#}X&]#}X&^#}X&`#}X~Oq#`O&a#aO~O&`ZOU$OXV$OXf$OXg$OXh$OXt$OXw$OXx$OXy$OXz$OX{$OX|$OX}$OX!O$OX!P$OX!Q$OX!R$OX$Y$OX$`$OX&Y$OX$h$OX$g$OX~O$f#dO~O$h#fO~PYOd#OOe#OOUmaVmafmagmahmatmawmaxmaymazma{ma|ma}ma!Oma!Pma!Qma!Rma$Yma$`ma&Yma$hma$gma~OP#iOQ#jOR#jOW$iPX$iPY$iPZ$iP[$iP]$iP^$iP_$iP`$iPa$iPb$iPc$iP!V$iP#z$iP$W$iP$a$iP$b$iP$f$iP$j$iP$k$iP$l$iP$m$iP$n$iP$o$iP$p$iP$q$iP$r$iP$s$iP$t$iP$u$iP$v$iP$w$iP$x$iP$y$iP$z$iP${$iP$|$iP$}$iP%O$iP%P$iP%Q$iP%R$iP%S$iP%T$iP%U$iP%V$iP%W$iP%X$iP%Y$iP%Z$iP%[$iP%]$iP%^$iP%_$iP%`$iP%a$iP%b$iP%c$iP%d$iP%e$iP%f$iP%g$iP%h$iP%i$iP%j$iP%k$iP%l$iP%m$iP%n$iP%o$iP%p$iP%q$iP%r$iP%s$iP%t$iP%u$iP%v$iP%w$iP%x$iP%y$iP%z$iP%{$iP%|$iP%}$iP&O$iP&P$iP&Q$iP&R$iP&S$iP&T$iP&U$iP&V$iP&W$iP&X$iP&[$iP&]$iP&^$iP&`$iP~O#v#rO~O{#tO#v#sO~Oi#vOj#vO$aVO$bVO&[#uO&]#uO&^VO~O$h#yO~P(OOt#SOU#|aV#|af#|ag#|ah#|aw#|ax#|ay#|az#|a{#|a|#|a}#|a!O#|a!P#|a!Q#|a!R#|a$Y#|a$`#|a&Y#|a$h#|a$g#|a~O!R#zO$S#zO$T#zO$U#zO$V#zO~O$g#|O&a$OO~Oq$QO$h$PO~O$g$RO$h$TO~P(OOQ#jOR#jOW$iXX$iXY$iXZ$iX[$iX]$iX^$iX_$iX`$iXa$iXb$iXc$iX!V$iX#z$iX$W$iX$a$iX$b$iX$f$iX$j$iX$k$iX$l$iX$m$iX$n$iX$o$iX$p$iX$q$iX$r$iX$s$iX$t$iX$u$iX$v$iX$w$iX$x$iX$y$iX$z$iX${$iX$|$iX$}$iX%O$iX%P$iX%Q$iX%R$iX%S$iX%T$iX%U$iX%V$iX%W$iX%X$iX%Y$iX%Z$iX%[$iX%]$iX%^$iX%_$iX%`$iX%a$iX%b$iX%c$iX%d$iX%e$iX%f$iX%g$iX%h$iX%i$iX%j$iX%k$iX%l$iX%m$iX%n$iX%o$iX%p$iX%q$iX%r$iX%s$iX%t$iX%u$iX%v$iX%w$iX%x$iX%y$iX%z$iX%{$iX%|$iX%}$iX&O$iX&P$iX&Q$iX&R$iX&S$iX&T$iX&U$iX&V$iX&W$iX&X$iX&[$iX&]$iX&^$iX&`$iX~O&Z$]O&_$^O~O#v$_O~O$a#]O$b#]O&^#]O~O$f$`O~O#z$aO$W$aO~Oq#`O&a$cO~O$g$dO&a$cO~O$g$fO$h$hO~O$g$RO$h$kO~OS$lOT$lOWvaXvaYvaZva[va]va^va_va`vaavabvacva!Vva#zva$Wva$ava$bva$fva$jva$kva$lva$mva$nva$ova$pva$qva$rva$sva$tva$uva$vva$wva$xva$yva$zva${va$|va$}va%Ova%Pva%Qva%Rva%Sva%Tva%Uva%Vva%Wva%Xva%Yva%Zva%[va%]va%^va%_va%`va%ava%bva%cva%dva%eva%fva%gva%hva%iva%jva%kva%lva%mva%nva%ova%pva%qva%rva%sva%tva%uva%vva%wva%xva%yva%zva%{va%|va%}va&Ova&Pva&Qva&Rva&Sva&Tva&Uva&Vva&Wva&Xva&[va&]va&^va&`va~Ot#SOUsqfsqgsqhsqzsq{sq|sq}sq!Osq!Psq!Qsq!Rsq$Ysq$`sq&Ysq$hsq$gsq~OVsqwsqxsqysq~PLTOV#SOw#SOx#SOy#SO~PLTOV#SOt#SOw#SOx#SOy#SOz#TO{#TOUsqfsqgsqhsq$Ysq$`sq&Ysq$hsq$gsq~O|sq}sq!Osq!Psq!Qsq!Rsq~PNPO|#UO}#UO!O#UO!P#UO!Q#UO!R#UO~PNPOV#SOf#VOh#VOt#SOw#SOx#SOy#SOz#TO{#TO|#UO}#UO!O#UO!P#UO!Q#UO!R#UO~OUsqgsq$Ysq$`sq&Ysq$hsq$gsq~P!!RO#v$nO&Z$mO~O$h$oO~Oq#`O&a$pO~Oq$rO$h$sO~O$g$tO$h$sO~O$g$^a$h$^a~P(OO$f#dOWviXviYviZvi[vi]vi^vi_vi`viavibvicvi!Vvi#zvi$Wvi$avi$bvi$jvi$kvi$lvi$mvi$nvi$ovi$pvi$qvi$rvi$svi$tvi$uvi$vvi$wvi$xvi$yvi$zvi${vi$|vi$}vi%Ovi%Pvi%Qvi%Rvi%Svi%Tvi%Uvi%Vvi%Wvi%Xvi%Yvi%Zvi%[vi%]vi%^vi%_vi%`vi%avi%bvi%cvi%dvi%evi%fvi%gvi%hvi%ivi%jvi%kvi%lvi%mvi%nvi%ovi%pvi%qvi%rvi%svi%tvi%uvi%vvi%wvi%xvi%yvi%zvi%{vi%|vi%}vi&Ovi&Pvi&Qvi&Rvi&Svi&Tvi&Uvi&Vvi&Wvi&Xvi&[vi&]vi&^vi&`vi~O&Z$wO~Oq#`O~Oq$rO$h$xO~Oq$rO~O',
  goto: ")T$jPPPPPPPPPPPPPPPPPPPPPPPPPPPPP$k$w%T%ZP%d$kP%m%tPPPPPPPPPPP$k&O&[P&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[&[$kP&h$k$kP$k$k&w$k'T'd'lPPPPP$kP'oP'r'x(OPPPPP(UPPP(te^OXY#P#m#n#o#p#q$ReROXY#P#m#n#o#p#q$RQ#QRR#h#RQ#e#OQ$U#jR$v$lQ#RRQ#[UR#h#QZ#l#S#T#U#V#WY#k#S#T#U#V#WR$V#leUOXY#P#m#n#o#p#q$ReTOXY#P#m#n#o#p#q$Rd^OXY#P#m#n#o#p#q$RR#w#ZeYOXY#P#m#n#o#p#q$Rd]OXY#P#m#n#o#p#q$RR#c[Q#bZV$b#|$d$qR#{#`R#x#ZQ$g$QR$u$gQ$S#gR$j$SQ#}#bR$e#}QSOQ#^XQ#_YQ#g#PQ$W#mQ$X#nQ$Y#oQ$Z#pQ$[#qR$i$RQ#m#SQ#n#TQ#o#UQ#p#VR#q#W",
  nodeNames:
    '⚠ Bool Ignoring On GroupLeft GroupRight Offset Atan2 Avg Bottomk Count CountValues Group Max Min Quantile Stddev Stdvar Sum Topk By Without And Or Unless Start End LineComment PromQL AggregateExpr AggregateOp AggregateModifier GroupingLabels LabelName FunctionCallBody BinaryExpr Pow BoolModifier MatchingModifierClause Mul Div Mod Add Sub Eql Gte Gtr Lte Lss Neq FunctionCall FunctionIdentifier AbsentOverTime Identifier Absent Abs Acos Acosh Asin Asinh Atan Atanh AvgOverTime Ceil Changes Clamp ClampMax ClampMin Cos Cosh CountOverTime DaysInMonth DayOfMonth DayOfWeek DayOfYear Deg Delta Deriv Exp Floor HistogramCount HistogramFraction HistogramQuantile HistogramStdDev HistogramStdVar HistogramSum HistogramAvg HoltWinters Hour Idelta Increase Irate LabelReplace LabelJoin LastOverTime Ln Log10 Log2 MadOverTime MaxOverTime MinOverTime Minute Month Pi PredictLinear PresentOverTime QuantileOverTime Rad Rate Resets Round Scalar Sgn Sin Sinh Sort SortDesc SortByLabel SortByLabelDesc Sqrt StddevOverTime StdvarOverTime SumOverTime Tan Tanh Timestamp Time Vector Year MatrixSelector Duration NumberLiteral OffsetExpr ParenExpr StringLiteral SubqueryExpr UnaryExpr UnaryOp VectorSelector LabelMatchers LabelMatcher MatchOp EqlSingle EqlRegex NeqRegex BindWithDashVar DashVariableLiteral StepInvariantExpr At AtModifierPreprocessors MetricName',
  maxTerm: 247,
  nodeProps: [['group', -13, 29, 35, 50, 129, 131, 132, 133, 134, 135, 136, 138, 146, 147, 'Expr']],
  propSources: [promQLHighLight, promQLIndent],
  skippedNodes: [0, 27],
  repeatNodeCount: 3,
  tokenData:
    "<`~RxX^#opq#oqr$drs$wst&otu'Wuv'iwx'nxy)ayz)fz{)k{|)p|})w}!O)|!O!P*T!P!Q+T!Q!R+Y!R![,o![!]8c!^!_9a!_!`9n!`!a:T!b!c:b!c!}:u!}#O;]#P#Q;b#Q#R;g#R#S:u#S#T;l#T#o:u#o#p<U#q#r<Z#y#z#o$f$g#o#BY#BZ#o$IS$I_#o$I|$JO#o$JT$JU#o$KV$KW#o&FU&FV#o~#tY$d~X^#opq#o#y#z#o$f$g#o#BY#BZ#o$IS$I_#o$I|$JO#o$JT$JU#o$KV$KW#o&FU&FV#o~$gQ!_!`$m#r#s$r~$rO!R~~$wO$U~~$|W#z~OY$wZr$wrs%fs#O$w#O#P%k#P;'S$w;'S;=`&i<%lO$w~%kO#z~~%nRO;'S$w;'S;=`%w;=`O$w~%|X#z~OY$wZr$wrs%fs#O$w#O#P%k#P;'S$w;'S;=`&i;=`<%l$w<%lO$w~&lP;=`<%l$w~&tSk~OY&oZ;'S&o;'S;=`'Q<%lO&o~'TP;=`<%l&o~']S$W~!Q!['W!c!}'W#R#S'W#T#o'W~'nOy~~'sW#z~OY'nZw'nwx%fx#O'n#O#P(]#P;'S'n;'S;=`)Z<%lO'n~(`RO;'S'n;'S;=`(i;=`O'n~(nX#z~OY'nZw'nwx%fx#O'n#O#P(]#P;'S'n;'S;=`)Z;=`<%l'n<%lO'n~)^P;=`<%l'n~)fO$f~~)kO$h~~)pOw~R)wO&]PzQ~)|O$g~R*TO&[P{QP*WP!Q![*ZP*`R&^P!Q![*Z!g!h*i#X#Y*iP*lR{|*u}!O*u!Q![*{P*xP!Q![*{P+QP&^P!Q![*{~+YOx~V+adqS&^P!O!P*Z!Q![,o!c!g.U!g!h.g!h!}.U#R#S.U#T#W.U#W#X/c#X#Y.g#Y#[.U#[#]0k#]#a.U#a#b1m#b#g.U#g#h3q#h#k.U#k#l4m#l#m7a#m#n5{#n#o.UV,vdqS&^P!O!P*Z!Q![,o!c!g.U!g!h.g!h!}.U#R#S.U#T#W.U#W#X/c#X#Y.g#Y#[.U#[#]0k#]#a.U#a#b1m#b#g.U#g#h3q#h#k.U#k#l4m#l#m.U#m#n5{#n#o.US.ZSqS!Q![.U!c!}.U#R#S.U#T#o.UT.lUqS{|*u}!O*u!Q![/O!c!}.U#R#S.U#T#o.UT/VSqS&^P!Q![/O!c!}.U#R#S.U#T#o.UU/jS#vQqS!Q![/v!c!}.U#R#S.U#T#o.UU/{YqS!Q![/v!c!}.U#R#S.U#T#[.U#[#]0k#]#a.U#a#b1m#b#g.U#g#h3q#h#o.UU0rS#vQqS!Q![1O!c!}.U#R#S.U#T#o.UU1TWqS!Q![1O!c!}.U#R#S.U#T#a.U#a#b1m#b#g.U#g#h3q#h#o.UU1tU#vQqS!Q![2W!c!}.U#R#S.U#T#g.U#g#h3^#h#o.UU2]WqS!Q![2W!c!}.U#R#S.U#T#a.U#a#b2u#b#g.U#g#h3q#h#o.UU2zUqS!Q![.U!c!}.U#R#S.U#T#g.U#g#h3^#h#o.UU3eS#vQqS!Q![.U!c!}.U#R#S.U#T#o.UU3xS#vQqS!Q![4U!c!}.U#R#S.U#T#o.UU4ZUqS!Q![4U!c!}.U#R#S.U#T#a.U#a#b2u#b#o.UU4tS#vQqS!Q![5Q!c!}.U#R#S.U#T#o.UU5V[qS!Q![5Q!c!}.U#R#S.U#T#W.U#W#X/c#X#[.U#[#]0k#]#a.U#a#b1m#b#g.U#g#h3q#h#o.UU6SS#vQqS!Q![6`!c!}.U#R#S.U#T#o.UU6e^qS!Q![6`!c!}.U#R#S.U#T#W.U#W#X/c#X#[.U#[#]0k#]#a.U#a#b1m#b#g.U#g#h3q#h#k.U#k#l4m#l#o.UT7fUqS!Q![7x!c!i7x!i!}.U#R#S.U#T#Z7x#Z#o.UT8PUqS&^P!Q![7x!c!i7x!i!}.U#R#S.U#T#Z7x#Z#o.U_8lT&_W$VS!VR!Q![8{![!]8{!c!}8{#R#S8{#T#o8{R9QT!VR!Q![8{![!]8{!c!}8{#R#S8{#T#o8{~9fP!Q~!_!`9i~9nO!P~~9sQ$SS!_!`9y#r#s:OQ:OO|Q~:TO$T~~:YP!O~!_!`:]~:bO}~U:iS$YQqS!Q![.U!c!}.U#R#S.U#T#o.UV:|T!VRqS!Q![:u![!]8{!c!}:u#R#S:u#T#o:u~;bO&Y~~;gO&Z~~;lOt~~;oTO#S;l#S#T%f#T;'S;l;'S;=`<O<%lO;l~<RP;=`<%l;l~<ZO&`~~<`O&a~",
  tokenizers: [0, 1, 2, 3],
  topRules: { PromQL: [0, 28], MetricName: [1, 150] },
  specialized: [
    {
      term: 53,
      get: (value, stack) => specializeIdentifier(value, stack) << 1,
      external: specializeIdentifier,
    },
    {
      term: 53,
      get: (value, stack) => (extendIdentifier(value, stack) << 1) | 1,
      external: extendIdentifier,
      extend: true,
    },
    { term: 53, get: (value) => spec_Identifier[value] || -1 },
  ],
  tokenPrec: 0,
});
// This file was generated by lezer-generator. You probably shouldn't edit it.
export const inf = 155,
  nan = 156,
  Bool = 1,
  Ignoring = 2,
  On = 3,
  GroupLeft = 4,
  GroupRight = 5,
  Offset = 6,
  Atan2 = 7,
  Avg = 8,
  Bottomk = 9,
  Count = 10,
  CountValues = 11,
  Group = 12,
  Max = 13,
  Min = 14,
  Quantile = 15,
  Stddev = 16,
  Stdvar = 17,
  Sum = 18,
  Topk = 19,
  By = 20,
  Without = 21,
  And = 22,
  Or = 23,
  Unless = 24,
  Start = 25,
  End = 26,
  LineComment = 27,
  PromQL = 28,
  AggregateExpr = 29,
  AggregateOp = 30,
  AggregateModifier = 31,
  GroupingLabels = 32,
  LabelName = 33,
  FunctionCallBody = 34,
  BinaryExpr = 35,
  Pow = 36,
  BoolModifier = 37,
  MatchingModifierClause = 38,
  Mul = 39,
  Div = 40,
  Mod = 41,
  Add = 42,
  Sub = 43,
  Eql = 44,
  Gte = 45,
  Gtr = 46,
  Lte = 47,
  Lss = 48,
  Neq = 49,
  FunctionCall = 50,
  FunctionIdentifier = 51,
  AbsentOverTime = 52,
  Identifier = 53,
  Absent = 54,
  Abs = 55,
  Acos = 56,
  Acosh = 57,
  Asin = 58,
  Asinh = 59,
  Atan = 60,
  Atanh = 61,
  AvgOverTime = 62,
  Ceil = 63,
  Changes = 64,
  Clamp = 65,
  ClampMax = 66,
  ClampMin = 67,
  Cos = 68,
  Cosh = 69,
  CountOverTime = 70,
  DaysInMonth = 71,
  DayOfMonth = 72,
  DayOfWeek = 73,
  DayOfYear = 74,
  Deg = 75,
  Delta = 76,
  Deriv = 77,
  Exp = 78,
  Floor = 79,
  HistogramCount = 80,
  HistogramFraction = 81,
  HistogramQuantile = 82,
  HistogramStdDev = 83,
  HistogramStdVar = 84,
  HistogramSum = 85,
  HistogramAvg = 86,
  HoltWinters = 87,
  Hour = 88,
  Idelta = 89,
  Increase = 90,
  Irate = 91,
  LabelReplace = 92,
  LabelJoin = 93,
  LastOverTime = 94,
  Ln = 95,
  Log10 = 96,
  Log2 = 97,
  MadOverTime = 98,
  MaxOverTime = 99,
  MinOverTime = 100,
  Minute = 101,
  Month = 102,
  Pi = 103,
  PredictLinear = 104,
  PresentOverTime = 105,
  QuantileOverTime = 106,
  Rad = 107,
  Rate = 108,
  Resets = 109,
  Round = 110,
  Scalar = 111,
  Sgn = 112,
  Sin = 113,
  Sinh = 114,
  Sort = 115,
  SortDesc = 116,
  SortByLabel = 117,
  SortByLabelDesc = 118,
  Sqrt = 119,
  StddevOverTime = 120,
  StdvarOverTime = 121,
  SumOverTime = 122,
  Tan = 123,
  Tanh = 124,
  Timestamp = 125,
  Time = 126,
  Vector = 127,
  Year = 128,
  MatrixSelector = 129,
  Duration = 130,
  NumberLiteral = 131,
  OffsetExpr = 132,
  ParenExpr = 133,
  StringLiteral = 134,
  SubqueryExpr = 135,
  UnaryExpr = 136,
  UnaryOp = 137,
  VectorSelector = 138,
  LabelMatchers = 139,
  LabelMatcher = 140,
  MatchOp = 141,
  EqlSingle = 142,
  EqlRegex = 143,
  NeqRegex = 144,
  BindWithDashVar = 145,
  DashVariableLiteral = 146,
  StepInvariantExpr = 147,
  At = 148,
  AtModifierPreprocessors = 149,
  MetricName = 150;

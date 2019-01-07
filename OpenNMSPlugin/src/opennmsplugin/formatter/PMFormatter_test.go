/*
* Copyright 2018 Nokia
* Licensed under BSD 3-Clause Clear License,
* see LICENSE file for details.
 */

package formatter

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"opennmsplugin/config"

	log "github.com/sirupsen/logrus"
)

var testPMData = `<?xml version="1.0" encoding="UTF-8"?>
<measCollecFile xmlns="http://www.3gpp.org/ftp/specs/archive/32_series/32.435#measCollec">
<fileHeader fileFormatVersion="32.435 V10.0" vendorName="Nokia Networks">
<fileSender elementType="elmnt_type"/>
<measCollec beginTime="2018-05-02T09:45:00+03:00"/>
</fileHeader>
<measData>
<managedElement/>
<measInfo measInfoId="meas_id">
<granPeriod duration="test" endTime="2018-05-02T10:00:00+03:00"/>
<repPeriod duration="test"/>
<measTypes>M0 M1 M2 M3 M4 M5 M6 M7 M8 M9 M10 M11 M12 M13 M14 M15 M16 M17 M18 M19 M20 M21 M22 M23 M24 M25 M26 M27 M28 M29 M30 M31 M32 M33 M34 M35 M36 M37 M38 M39 M40 M41 M42 M43 M44 M45 M46 M47 M48 M49 M50 M51 M52 M53 M54 M55 M56 M57 M58 M59 M60 M61 M62 M63 M64 M65 M66 M67 M68 M69 M70 M71 M72 M73 M74 M75 M76 M77 M78 M79 M80 M81 M82 M83 M84 M85 M86 M87 M88 M89 M90 M91 M92 M93 M94 M95 M96 M97 M98 M99 M100 M101 M102 M103 M104 M105 M106 M107 M108 M109 M110 M111 M112 M113 M114 M115 M116 M117 M118 M119 M120 M121 M122 M123 M124 M125 M126 M127 M128 M129 M130 M131 M132 M133 M134 M135 M136 M137 M138 M139 M140 M141 M142 M143 M144 M145 M146 M147 M148 M149 M150 M151 M152 M153 M154 M155 M156 M157 M158 M159 M160 M161 M162 M163 M164 M165 M166 M167 M168 M169 M170 M171 M172 M173 M174 M175 M176 M177 M178 M179 M180 M181 M182 M183 M184 M185 M186 M187 M188 M189 M190 M191 M192 M193 M194 M195 M196 M197 M198 M199 M200 M201 M202 M203 M204 M205 M206 M207 M208 M209 M210 M211 M212 M213 M214 M215 M216 M217 M218 M219 M220 M221 M222 M223 M224 M225 M226 M227 M228 M229 M230 M231 M232 M233 M234 M235 M236 M237 M238 M239 M240 M241 M242 M243 M244 M245 M246 M247 M248 M249 M250 M251 M252 M253 M254 M255 M256 M257 M258 M259 M260 M261 M262 M263 M264 M265 M266 M267 M268 M269 M270 M271 M272 M273 M274 M275 M276 M277 M278 M279 M280 M281 M282 M283 M284 M285 M286 M287 M288 M289 M290 M291 M292 M293 M294 M295 M296 M297 M298 M299 M300 M301 M302 M303 M304 M305 M306 M307 M308 M309 M310 M311 M312 M313 M314 M315 M316 M317 M318 M319 M320 M321 M322 M323 M324 M325 M326 M327 M328 M329 M330 M331 M332 M333 M334 M335 M336 M337 M338 M339 M340 M341 M342 M343 M344 M345 M346 M347 M348 M349 M350 M351 M352 M353 M354 M355 M356 M357 M358 M359 M360 M361 M362 M363 M364 M365 M366 M367 M368 M369 M370 M371 M372 M373 M374 M375 M376 M377 M378 M379 M380 M381 M382 M383 M384 M385 M386 M387 M388 M389 M390 M391 M392 M393 M394 M395 M396 M397 M398 M399 M400 M401 M402 M403 M404 M405 M406 M407 M408 M409 M410 M411 M412 M413 M414 M415 M416 M417 M418 M419 M420 M421 M422 M423 M424 M425 M426 M427 M428 M429 M430 M431 M432 M433 M434 M435 M436 M437 M438 M439 M440 M441 M442 M443 M444 M445 M446 M447 M448 M449 M450 M451 M452 M453 M454 M455 M456 M457 M458 M459 M460 M461 M462 M463 M464 M465 M466 M467 M468 M469 M470 M471 M472 M473 M474 M475 M476 M477 M478 M479 M480 M481 M482 M483 M484 M485 M486 M487 M488 M489 M490 M491 M492 M493 M494 M495 M496 M497 M498 M499 M500 M501 M502 M503 M504 M505 M506 M507 M508 M509 M510 M511 M512 M513 M514 M515 M516 M517 M518 M519 M520 M521 M522 M523 M524 M525 M526 M527 M528 M529 M530 M531 M532 M533 M534 M535 M536 M537 M538 M539 M540 M541 M542 M543 M544 M545 M546 M547 M548 M549 M550 M551 M552 M553 M554 M555 M556 M557 M558 M559</measTypes>
<measValue measObjLdn="test_dn">
<measResults>0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0</measResults>
</measValue>
</measInfo>
</measData>
<fileFooter>
<measCollec endTime="2018-05-02T10:00:00+03:00"/>
</fileFooter>
</measCollecFile>`

func TestRenameFile(t *testing.T) {
	fileName := renameFile("PM201805021002+030024LNBTS_-_39.xml", "12345")
	if fileName == "" {
		t.Fail()
	}
}

func TestRenameFileWithWrongFormat(t *testing.T) {
	fileName := renameFile("test_file", "12345")
	if fileName != "" {
		t.Fail()
	}
}

func TestFormatPMData(t *testing.T) {
	pmConfig := config.PMConfig{
		DestinationDir: "./tmp",
		ForeignID:      "12345",
	}
	fileName := "./PM201805021002+030024LNBTS_-_39.xml"
	renamedFile := renameFile(fileName, pmConfig.ForeignID)

	err := createTestData(fileName, testPMData)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(fileName)

	err = os.MkdirAll(pmConfig.DestinationDir, os.ModePerm)
	if err != nil {
		t.Error(err)
	}

	FormatPMData(fileName, pmConfig)
	defer os.RemoveAll(pmConfig.DestinationDir)
	fileContent, err := ioutil.ReadFile("./tmp/" + renamedFile)
	if err != nil || len(fileContent) == 0 {
		t.Fail()
	}
}

func TestFormatPMDataWithNonExistingDir(t *testing.T) {
	pmConfig := config.PMConfig{
		DestinationDir: "./tmp",
		ForeignID:      "12345",
	}
	fileName := "./PM201805021002+030024LNBTS_-_39.xml"
	err := createTestData(fileName, testPMData)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(fileName)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	FormatPMData(fileName, pmConfig)
	if !strings.Contains(buf.String(), "no such file or directory") {
		t.Fail()
	}
}

func TestFormatPMDataWithNonExistingFile(t *testing.T) {
	pmConfig := config.PMConfig{
		DestinationDir: "./tmp",
		ForeignID:      "12345",
	}
	if _, err := os.Stat(pmConfig.DestinationDir); os.IsNotExist(err) {
		os.MkdirAll(pmConfig.DestinationDir, os.ModePerm)
	}
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	FormatPMData("test_file.xml", pmConfig)
	defer os.RemoveAll(pmConfig.DestinationDir)
	if !strings.Contains(buf.String(), "no such file or directory") {
		t.Fail()
	}
}

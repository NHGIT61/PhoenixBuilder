package I18n

var I18nDict_en_UK map[uint16]string = map[uint16]string{
	ACME_FailedToGetCommand:             "Failed to get acme command.",
	ACME_FailedToSeek:                   "Invalid acme file since seek was failed.",
	ACME_StructureErrorNotice:           "Invalid structure",
	ACME_UnknownCommand:                 "Unknown ACME command",
	Auth_BackendError:                   "Backend Error",
	Auth_FailedToRequestEntry:           "Failed to request entry for your server, please cheque whether the password is correct and please turn off the level limitation",
	Auth_HelperNotCreated:               "Helper user haven't been created, please go create it on FastBuilder User Centre.",
	Auth_InvalidFBVersion:               "Invalid FastBuilder version, please update.",
	Auth_InvalidHelperUsername:          "Invalid username for helper user, please configure it on FastBuilder User Centre.",
	Auth_InvalidToken:                   "Invalid login token.",
	Auth_InvalidUser:                    "Invalid user for FastBuilder User Centre",
	Auth_ServerNotFound:                 "Server not found, please cheque your server's public state",
	Auth_UnauthorizedRentalServerNumber: "Unauthorised rental server number, please add it on your FastBuilder User Centre.",
	Auth_UserCombined:                   "Given user has been combined to another account, please login using another account's information.",
	Auth_FailedToRequestEntry_TryAgain:  "Failed to request server entry, please try again later.",
	BDump_Author:                        "Author",
	BDump_EarlyEOFRightWhenOpening:      "Failed to read file, early EOF? File may be corrupted",
	BDump_FailedToGetCmd1:               "Failed to get any argument for cmd[pos:0], file may corrupted",
	BDump_FailedToGetCmd2:               "Failed to get any argument for cmd[pos1], file may corrupted",
	BDump_FailedToGetCmd4:               "Failed to get any argument for cmd[pos2], file may corrupted",
	BDump_FailedToGetCmd6:               "Failed to get any argument for cmd[pos3], file may corrupted",
	BDump_FailedToGetCmd7_0:             "Failed to get any argument for cmd[pos4], file may corrupted",
	BDump_FailedToGetCmd7_1:             "Failed to get any argument for cmd[pos5], file may corrupted",
	BDump_FailedToGetCmd10:              "Failed to get any argument for cmd[pos6], file may corrupted",
	BDump_FailedToGetCmd11:              "Failed to get any argument for cmd[pos7], file may corrupted",
	BDump_FailedToGetCmd12:              "Failed to get any argument for cmd[pos8], file may corrupted",
	BDump_FailedToGetConstructCmd:       "Failed to get construction commands, file may corrupted",
	BDump_FailedToReadAuthorInfo:        "Failed to read author info, file may be corrupted",
	BDump_FileNotSigned:                 "File is not signed",
	BDump_FileSigned:                    "File is signed, signer: %s",
	BDump_NotBDX_Invheader:              "Not a bdx file (Invalid file header)",
	BDump_NotBDX_Invinnerheader:         "Not a bdx file (Invalid inner file header)",
	BDump_SignedVerifying:               "File is signed, verifying...",
	BDump_VerificationFailedFor:         "Failed to verify the file's signature due to: %v",
	BDump_Warn_Reserved:                 "WARNING: BDump/Import: Use of reserved command\n",
	CommandNotFound:                     "Command not found.",
	ConnectionEstablished:               "Successfully created minecraft dialler.",
	Copyright_Notice_Bouldev:            "Copyright (c) FastBuilder DevGroup, Bouldev 2022",
	Copyright_Notice_Contrib:            "Contributors: Ruphane, CAIMEO, CMA2401PT",
	Crashed_No_Connection:               "connection haven't got established after significant amount of time",
	Crashed_OS_Windows:                  "Press ENTER to exit.",
	Crashed_StackDump_And_Error:         "Stack dump has been shown above, error:",
	Crashed_Tip:                         "Oh no! FastBuilder Phoenix crashed!",
	CurrentDefaultDelayMode:             "Current default delay mode",
	CurrentTasks:                        "Current tasks:",
	DelayModeSet:                        "Delay mode set",
	DelayModeSet_DelayAuto:              "Delay has been automatically set to: %d",
	DelayModeSet_ThresholdAuto:          "Delay threshold has been automatically set to: %d",
	DelaySet:                            "Delay set",
	DelaySetUnavailableUnderNoneMode:    "[delay set] is unavailable with delay mode: none",
	DelayThreshold_OnlyDiscrete:         "Delay threshold is only available with delay mode: discrete",
	DelayThreshold_Set:                  "Delay threshold has been set to: %d",
	ERRORStr:                            "ERROR",
	EnterPasswordForFBUC:                "Enter your password for FastBuilder User Centre: ",
	Enter_FBUC_Username:                 "Enter your FastBuilder User Centre username: ",
	Enter_Rental_Server_Code:            "Please enter your rental server number: ",
	Enter_Rental_Server_Password:        "Enter Password (Press [Enter] if not set, input won't be echoed): ",
	ErrorIgnored:                        "Error ignored.",
	Error_MapY_Exceed:                   "In 3DMap, MapY should be in range of [20~255] (Your Input = %v)",
	FBUC_LoginFailed:                    "Incorrect username or password",
	FBUC_Token_ErrOnCreate:              "Error creating token file: ",
	FBUC_Token_ErrOnGen:                 "Failed to generate temp token",
	FBUC_Token_ErrOnRemove:              "Failed to remove token file: %v",
	FBUC_Token_ErrOnSave:                "Error saving token: ",
	FileCorruptedError:                  "File corrupted",
	Get_Warning:                         "",
	IgnoredStr:                          "ignored",
	InvalidFileError:                    "Invalid file",
	InvalidPosition:                     "No position got. (ignorable)",
	Lang_Config_ErrOnCreate:             "Error creating language config file: %v",
	Lang_Config_ErrOnSave:               "Error saving language config: %v",
	LanguageName:                        "English (UK)",
	LanguageUpdated:                     "Language preference has been updated",
	Logout_Done:                         "Logged out from FastBuilder User Centre.",
	Menu_BackButton:                     "< Back",
	Menu_Cancel:                         "Cancel",
	Menu_CurrentPath:                    "Current path",
	Menu_ExcludeCommandsOption:          "Exclude Commands",
	Menu_GetEndPos:                      "getEndPos",
	Menu_GetPos:                         "getPos",
	Menu_InvalidateCommandsOption:       "Invalidate Commands",
	Menu_Quit:                           "Quit Programme",
	Menu_StrictModeOption:               "Strict Mode",
	NotAnACMEFile:                       "Invalid file, not an ACME structure.",
	Notice_iSH_Location_Service:         "You are using iSH simulator, location service is required for foreground, no location data will be saved or used. You can terminate it anytime.",
	Notice_CheckUpdate:                  "Checking update, please wait...",
	Notice_OK:                           "OK\n",
	Notice_UpdateAvailable:              "A newer version (%s) of PhoenixBuilder is available.\n",
	Notice_UpdateNotice:                 "Please update.\n",
	Notice_ZLIB_CVE:                     "Your zlib version (%s) is too old since it contains confirmed CVE vulnerability, updating suggested",
	Notify_NeedOp:                       "FastBuilder requires operator privilege.",
	Notify_TurnOnCmdFeedBack:            "FastBuilder requires gamerule sendcommandfeedback to be true, we have already turn it on, and remember to turn it off",
	Omega_Enabled:                       "Omega System Enabled!",
	Parsing_UnterminatedEscape:          "Unterminated escape",
	Parsing_UnterminatedQuotedString:    "Unterminated quoted string",
	PositionGot:                         "Position got",
	PositionGot_End:                     "End Position got",
	PositionSet:                         "Position set",
	PositionSet_End:                     "End position set",
	QuitCorrectly:                       "Quit correctly",
	Sch_FailedToResolve:                 "Failed to resolve file",
	SelectLanguageOnConsole:             "Please select your new preferred language on console.",
	ServerCodeTrans:                     "Server",
	SimpleParser_Int_ParsingFailed:      "Parser: failed to parse an int argument",
	SimpleParser_InvEnum:                "Parser: Invalid enum value, allowed values are: %s.",
	SimpleParser_Invalid_decider:        "Parser: Invalid decider",
	SimpleParser_Too_few_args:           "Parser: Too few arguments",
	Special_Startup:                     "Enabled language: English (UK)\n",
	TaskCreated:                         "Task Created",
	TaskDisplayModeSet:                  "Task status display mode has been set to: %s.",
	TaskFailedToParseCommand:            "Failed to parse command: %v",
	TaskNotFoundMessage:                 "Could not find a valid task by provided task id.",
	TaskPausedNotice:                    "[Task %d] - Paused",
	TaskResumedNotice:                   "[Task %d] - Resumed",
	TaskStateLine:                       "ID %d - CommandLine:\"%s\", State: %s, Delay: %d, DelayMode: %s, DelayThreshold: %d",
	TaskStoppedNotice:                   "[Task %d] - Stopped",
	TaskTTeIuKoto:                       "Task",
	TaskTotalCount:                      "Total: %d",
	TaskTypeCalculating:                 "Calculating",
	TaskTypeDied:                        "Died",
	TaskTypePaused:                      "Paused",
	TaskTypeRunning:                     "Running",
	TaskTypeSpecialTaskBreaking:         "SpecialTask:Breaking",
	TaskTypeSwitchedTo:                  "Task creation type set to: %s.",
	TaskTypeUnknown:                     "Unknown",
	Task_D_NothingGenerated:             "[Task %d] Nothing generated.",
	Task_DelaySet:                       "[Task %d] - Delay set: %d",
	Task_ResumeBuildFrom:                "Resume Build From Block Number %v ",
	Task_SetDelay_Unavailable:           "[setdelay] is unavailable with delay mode: none",
	Task_Summary_1:                      "[Task %d] %v block(s) have been changed.",
	Task_Summary_2:                      "[Task %d] Time used: %v second(s)",
	Task_Summary_3:                      "[Task %d] Average speed: %v blocks/second",
	UnsupportedACMEVersion:              "Unsupported ACME structure version. Only acme file version 1.2 is supported.",
	Warning_ACME_Deprecated:             "WARNING - `acme' is deprecated and will be removed in future, please migrate to other format instead.\n"
	Warning_UserHomeDir:                 "WARNING - Failed to obtain the user's home directory. made homedir=\".\";\n",
}

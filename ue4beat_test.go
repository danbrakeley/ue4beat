package main

import (
	"testing"
	"time"
)

type testCase struct {
	In        string
	Timestamp string
	Frame     int
	Category  string
	Message   string
	Level     string
}

var realCases = []testCase{
	{
		`sh: xdg-user-dir: command not found`,
		"", -1, "", `sh: xdg-user-dir: command not found`, LevelInfo,
	}, {
		`[S_API FAIL] SteamAPI_Init() failed; SteamAPI_IsSteamRunning() failed.`,
		"", -1, "", `[S_API FAIL] SteamAPI_Init() failed; SteamAPI_IsSteamRunning() failed.`, LevelInfo,
	}, {
		`CAppInfoCacheReadFromDiskThread took 0 milliseconds to initialize`,
		"", -1, "", `CAppInfoCacheReadFromDiskThread took 0 milliseconds to initialize`, LevelInfo,
	}, {
		`CApplicationManagerPopulateThread took 0 milliseconds to initialize (will have waited on CAppInfoCacheReadFromDiskThread)`,
		"", -1, "", `CApplicationManagerPopulateThread took 0 milliseconds to initialize (will have waited on CAppInfoCacheReadFromDiskThread)`, LevelInfo,
	}, {
		`Setting breakpad minidump AppID = 803370`,
		"", -1, "", `Setting breakpad minidump AppID = 803370`, LevelInfo,
	}, {
		`Increasing per-process limit of core file size to infinity.`,
		"", -1, "", `Increasing per-process limit of core file size to infinity.`, LevelInfo,
	}, {
		`LogPlatformFile: Using cached read wrapper`,
		"", -1, "LogPlatformFile", `Using cached read wrapper`, LevelInfo,
	}, {
		`LogStreaming: Display: Took  0.074s to configure plugins.`,
		"", -1, "LogStreaming", `Took  0.074s to configure plugins.`, LevelInfo,
	}, {
		`[33mLogFileManager: Warning: ReadFile failed: Count=0 BufferCount=22 Error=errno=21 (Is a directory)`,
		"", -1, "LogFileManager", `ReadFile failed: Count=0 BufferCount=22 Error=errno=21 (Is a directory)`, LevelWarning,
	}, {
		`[0mLogInit: Using libcurl 7.57.0`,
		"", -1, "LogInit", `Using libcurl 7.57.0`, LevelInfo,
	}, {
		`LogInit:  - built for x86_64-pc-linux-gnu`,
		"", -1, "LogInit", `- built for x86_64-pc-linux-gnu`, LevelInfo,
	}, {
		`LogInit:  - supports SSL with OpenSSL/1.0.2h`,
		"", -1, "LogInit", `- supports SSL with OpenSSL/1.0.2h`, LevelInfo,
	}, {
		`LogStreaming: Display: Async Loading initialized: Event Driven Loader: true, Async Loading Thread: false`,
		"", -1, "LogStreaming", `Async Loading initialized: Event Driven Loader: true, Async Loading Thread: false`, LevelInfo,
	}, {
		`LogInit: Object subsystem initialized`,
		"", -1, "LogInit", `Object subsystem initialized`, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:212][  0]LogLinux: Selected Device Profile: [LinuxServer]`,
		"2018-09-21T21:44:44.212Z", 0, "LogLinux", `Selected Device Profile: [LinuxServer]`, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:212][  0]LogInit: Applying CVar settings loaded from the selected device profile: [LinuxServer]`,
		"2018-09-21T21:44:44.212Z", 0, "LogInit", `Applying CVar settings loaded from the selected device profile: [LinuxServer]`, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:218][  0]LogInit: Linux hardware info:`,
		"2018-09-21T21:44:44.218Z", 0, "LogInit", `Linux hardware info:`, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:218][  0]LogInit:  - we are the first instance of this executable`,
		"2018-09-21T21:44:44.218Z", 0, "LogInit", `- we are the first instance of this executable`, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:218][  0]LogInit:  - this process' id (pid) is 15032, parent process' id (ppid) is 1`,
		"2018-09-21T21:44:44.218Z", 0, "LogInit", `- this process' id (pid) is 15032, parent process' id (ppid) is 1`, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:218][  0]LogInit:  - machine network name is 'ip-10-0-90-68.ec2.internal'`,
		"2018-09-21T21:44:44.218Z", 0, "LogInit", `- machine network name is 'ip-10-0-90-68.ec2.internal'`, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:950][  0]LogSynthBenchmark: Display:   Device Revision: 0x0`,
		"2018-09-21T21:44:44.95Z", 0, "LogSynthBenchmark", `Device Revision: 0x0`, LevelInfo,
	}, {
		`[33m[2018.09.21-21.44.44:950][  0]LogSynthBenchmark: Warning: RendererGPUBenchmark failed, look for "GPU Timing Frequency" in the log`,
		"2018-09-21T21:44:44.95Z", 0, "LogSynthBenchmark", `RendererGPUBenchmark failed, look for "GPU Timing Frequency" in the log`, LevelWarning,
	}, {
		`[0m[33m[2018.09.21-21.44.44:950][  0]LogSynthBenchmark: Warning: RendererGPUBenchmark failed, look for "GPU Timing Frequency" in the log`,
		"2018-09-21T21:44:44.95Z", 0, "LogSynthBenchmark", `RendererGPUBenchmark failed, look for "GPU Timing Frequency" in the log`, LevelWarning,
	}, {
		`[0m[2018.09.21-21.44.44:950][  0]LogSynthBenchmark: Display:   CPUIndex: 129.4`,
		"2018-09-21T21:44:44.95Z", 0, "LogSynthBenchmark", `CPUIndex: 129.4`, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:950][  0]LogSynthBenchmark: Display: `,
		"2018-09-21T21:44:44.95Z", 0, "LogSynthBenchmark", ``, LevelInfo,
	}, {
		`[2018.09.21-21.44.44:950][  0]LogSynthBenchmark: Display:          ... Total Time: 0.427421 sec`,
		"2018-09-21T21:44:44.95Z", 0, "LogSynthBenchmark", `... Total Time: 0.427421 sec`, LevelInfo,
	}, {
		`[31m[2018.09.21-21.44.45:273][  0]HaxeLog: Error: Runner.hx:234: The field is not optional but it was not set,{ Field => JwtSharedSecret }`,
		"2018-09-21T21:44:45.273Z", 0, "HaxeLog", `Runner.hx:234: The field is not optional but it was not set,{ Field => JwtSharedSecret }`, LevelError,
	}, {
		`[0m[2018.09.21-21.44.45:273][  0]LogInit: Display: Game Engine Initialized.`,
		"2018-09-21T21:44:45.273Z", 0, "LogInit", `Game Engine Initialized.`, LevelInfo,
	}, {
		`[2018.09.21-21.44.45:273][  0]LogGameplayTags: Display: UGameplayTagsManager::DoneAddingNativeTags. DelegateIsBound: 0`,
		"2018-09-21T21:44:45.273Z", 0, "LogGameplayTags", `UGameplayTagsManager::DoneAddingNativeTags. DelegateIsBound: 0`, LevelInfo,
	}, {
		`[2018.09.21-21.44.45:274][ 10]LogStats: UGameplayTagsManager::ConstructGameplayTagTree: Construct from data asset -  0.000 s`,
		"2018-09-21T21:44:45.274Z", 10, "LogStats", `UGameplayTagsManager::ConstructGameplayTagTree: Construct from data asset -  0.000 s`, LevelInfo,
	}, {
		`[2018.09.21-21.45.03:660][283]LogLevelStreaming: Verbose: Level /Game/Maps/Lobby:-1 CurrentState MakingVisible -> LoadedVisible (TargetState=LoadedVisible)`,
		"2018-09-21T21:45:03.66Z", 283, "LogLevelStreaming", `Level /Game/Maps/Lobby:-1 CurrentState MakingVisible -> LoadedVisible (TargetState=LoadedVisible)`, LevelInfo,
	},
}

var malformedCases = []testCase{
	{`Lo: Foo`, "", -1, "", `Lo: Foo`, LevelInfo},
	{`Log: Foo`, "", -1, "Log", `Foo`, LevelInfo},
	{` Log: Foo`, "", -1, "", `Log: Foo`, LevelInfo},
	{`LogTest:`, "", -1, "", `LogTest:`, LevelInfo},
	{`LogTest: `, "", -1, "LogTest", ``, LevelInfo},
	{`LogTest: Error:`, "", -1, "LogTest", `Error:`, LevelInfo},
	{`LogTest: Error: `, "", -1, "LogTest", ``, LevelError},
	{`LogTest: Error:  `, "", -1, "LogTest", ``, LevelError},
	{`: Error:`, "", -1, "", `: Error:`, LevelInfo},
	{`[33`, "", -1, "", `[33`, LevelInfo},
	{`[2018.09.21-21.44.44:950]Warning: LogBar: frank`, "2018-09-21T21:44:44.95Z", -1, "Warning", `LogBar: frank`, LevelInfo},
	{`[  a]Warning: LogBar: frank`, "", -1, "", `[  a]Warning: LogBar: frank`, LevelInfo},
	{`[  1a]Warning: LogBar: frank`, "", -1, "", `[  1a]Warning: LogBar: frank`, LevelInfo},
}

func runTestCases(t *testing.T, cases []testCase, fnParseLine func(string) UE4Line) {
	for i, c := range cases {
		u := fnParseLine(c.In)

		if len(c.Timestamp) == 0 {
			if u.Timestamp != nil {
				t.Errorf("Case %d: Expected nil Timestamp, actual is \"%v\"", i, u.Timestamp)
			}
		} else {
			timestamp, err := time.Parse(time.RFC3339, c.Timestamp)
			if err != nil {
				t.Fatalf("Case %d: Bad test case timestamp \"%s\": %v", i, c.Timestamp, err)
			}
			if u.Timestamp == nil || *u.Timestamp != timestamp {
				t.Errorf("Case %d: Expected %v, actual is %v", i, timestamp, u.Timestamp)
			}
		}
		if u.Frame != c.Frame {
			t.Errorf("Case %d: Expected Frame %d, actual is %d", i, c.Frame, u.Frame)
		}
		if u.Category != c.Category {
			t.Errorf("Case %d: Expected Category \"%s\", actual is \"%s\"", i, c.Category, u.Category)
		}
		if u.Message != c.Message {
			t.Errorf("Case %d: Expected Message \"%s\", actual is \"%s\"", i, c.Message, u.Message)
		}
		if u.Level != c.Level {
			t.Errorf("Case %d: Expected Level %v, actual is %v", i, c.Level, u.Level)
		}
	}
}

func TestParseLine(t *testing.T) {
	runTestCases(t, realCases, ParseLine)
	runTestCases(t, malformedCases, ParseLine)
}

func BenchmarkParseLine(b *testing.B) {
	// run ParseLine function b.N times
	for n := 0; n < b.N; n++ {
		ParseLine(realCases[n%len(realCases)].In)
	}
}

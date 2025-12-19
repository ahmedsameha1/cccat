exception NoArgs

fun catWithoutNumbers(streams) =
    case (TextIO.endOfStream(streams), streams) of
        (true, _) => print("")
        | (false, streams) => 
            (
                (case TextIO.input1(streams) of
                    SOME c => print(String.str(c))
                    | NONE => print("\n"));
             catWithoutNumbers(streams)
            )

fun catWithNumbers(stream, acc, acc2) =
    case (TextIO.endOfStream(stream), TextIO.input1(stream), TextIO.lookahead(stream), acc, acc2) of
        (true, _, _, _, false) => (TextIO.closeIn(stream); print("\n"))
        | (true, SOME c, SOME d, accc, true) =>
                    if String.str(c) = "\n"  
                    then (TextIO.closeIn(stream); print(String.concat([Int.toString(acc), " ", String.str(c)])))
                    else (TextIO.closeIn(stream); print(String.str(c)))
        | (true, SOME c, NONE, accc, true) =>
                    (TextIO.closeIn(stream); print(String.str(c)))
        | (false, SOME c, SOME d, accc, true) =>
                    if String.str(c) = "\n"  
                    then (print(String.concat([String.str(c), Int.toString(acc), " "])); catWithNumbers(stream, acc + 1, true))
                    else (print(String.str(c)); catWithNumbers(stream, acc, true))
        | (false, SOME c, SOME d, accc, false) =>
                    (print(String.concat([Int.toString(acc), " ", String.str(c)])); catWithNumbers(stream, acc + 1, true))
        | (false, SOME c, NONE, accc, true) => (print(String.str(c));catWithNumbers(stream, acc, true))
                    | _ => print("")

fun catWithNumbersNotEmptyLines(stream, acc, acc2) =
    case (TextIO.endOfStream(stream), TextIO.input1(stream), TextIO.lookahead(stream), acc, acc2) of
        (true, _, _, _, false) => (TextIO.closeIn(stream); print("\n"))
        | (true, SOME c, SOME d, accc, true) =>
                    if String.str(c) = "\n" andalso String.str(d) = "\n"  
                    then (TextIO.closeIn(stream); print(String.str(c)); print(String.str(d)))
                    else if String.str(c) = "\n" 
                    then (TextIO.closeIn(stream); print(String.concat([Int.toString(acc), " ", String.str(c), String.str(d)])))
                    else (TextIO.closeIn(stream); print(String.str(c)); print(String.str(d)))
        | (true, SOME c, NONE, accc, true) =>
                    (TextIO.closeIn(stream); print(String.str(c)))
        | (false, SOME c, SOME d, accc, true) =>
                    if String.str(c) = "\n" andalso String.str(d) = "\n"  
                    then (print(String.str(c)); catWithNumbersNotEmptyLines(stream, acc, true))
                    else if String.str(c) = "\n"  
                    then (print(String.concat([String.str(c), Int.toString(acc), " "])); catWithNumbersNotEmptyLines(stream, acc + 1, true))
                    else (print(String.str(c)); catWithNumbersNotEmptyLines(stream, acc, true))
        | (false, SOME c, SOME d, accc, false) =>
                    (print(String.concat([Int.toString(acc), " ", String.str(c)])); catWithNumbersNotEmptyLines(stream, acc + 1, true))
        | (false, SOME c, NONE, accc, true) => (print(String.str(c));catWithNumbersNotEmptyLines(stream, acc, true))
                    | _ => print("")

fun cat() = 
    let
    val arguments =  CommandLine.arguments ()
        val (streams, withNumbers, withNumbersNoEmptyLines) = case arguments of
            [] => ([TextIO.stdIn], false, false)
            |"-n"::[] => ([TextIO.stdIn], true, false)
            |"-nb"::[] => ([TextIO.stdIn], true, true)
            |"-bn"::[] => ([TextIO.stdIn], true, true)
            |"-b"::[] => ([TextIO.stdIn], true, true)
            |"-"::[] => ([TextIO.stdIn], false, false)
            |x::[] => ([TextIO.openIn x], false, false)
            |"-n"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true, false)
            |"-bn"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true, true)
            |"-nb"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true, true)
            |"-b"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true, true)
            |x => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) x), false, false)
    in
        if withNumbers andalso withNumbersNoEmptyLines
            then List.app (fn (x) => catWithNumbersNotEmptyLines(x, 1, false)) streams
            else if withNumbers
            then List.app (fn (x) => catWithNumbers(x, 1, false)) streams
            else List.app catWithoutNumbers streams
    end

fun main () = cat()
val _ = main ();
